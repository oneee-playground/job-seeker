package wanted

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/oneee-playground/job-seeker/search"
)

// TODO: make it configurable.
const baseURL = "https://www.wanted.co.kr"

type Platform struct {
	httpClient *http.Client
}

func NewPlatform(httpClient *http.Client) *Platform {
	return &Platform{httpClient: httpClient}
}

var _ search.Platform = (*Platform)(nil)

func (p *Platform) Search(ctx context.Context, opts search.Options) (<-chan search.Result, <-chan error) {
	results := make(chan search.Result, 10)
	errchan := make(chan error)

	go func() {
		defer close(errchan)
		defer close(results)

		found := 0

		const uriBase = "/api/chaos/navigation/v1/results?job_group_id=518&country=kr&job_sort=job.latest_order&locations=all&limit=30"

		u, _ := url.Parse(uriBase)
		query := u.Query()

		query.Add("years", strconv.Itoa(opts.ExpYears.Start()))
		query.Add("years", strconv.Itoa(opts.ExpYears.End()))

		u.RawQuery = query.Encode()

		nextURI := u.String()

		for nextURI != "" && (opts.Limit < 0 || found < opts.Limit) {
			list, err := p.fetchJobList(baseURL + nextURI)
			if err != nil {
				errchan <- err
				return
			}

			for _, job := range list.Data {
				detailURL := makeJobDetailURL(job.ID)

				detail, err := p.fetchJobDetail(detailURL)
				if err != nil {
					errchan <- err
					return
				}

				if ok := filterJob(detail.Job, opts); ok {
					expYears := search.ExpYears{detail.Job.AnnualFrom, detail.Job.AnnualTo}
					if expYears.End() == 100 {
						// expYears limit of wanted is 100. meaning it has no end limit.
						expYears.SetNoEndLimit()
					}

					results <- search.Result{
						Platform: search.PlatformWanted,
						Company:  detail.Job.Company.Name,
						Position: detail.Job.Detail.Position,
						ExpYears: expYears,
						URL:      makeJobDetailFrontendURL(job.ID),
					}
					found++
				}
			}

			if next := list.Links.Next; next != nil {
				nextURI = *next
			} else {
				nextURI = ""
			}
		}

	}()

	return results, errchan
}

func (p *Platform) fetchJobList(url string) (listResponse, error) {
	res, err := p.doFetch(url)
	if err != nil {
		return listResponse{}, err
	}

	var decoded listResponse
	if err := json.NewDecoder(res.Body).Decode(&decoded); err != nil {
		return listResponse{}, err
	}

	return decoded, nil
}

func (p *Platform) fetchJobDetail(url string) (detailResponse, error) {
	res, err := p.doFetch(url)
	if err != nil {
		return detailResponse{}, err
	}

	var decoded detailResponse
	if err := json.NewDecoder(res.Body).Decode(&decoded); err != nil {
		return detailResponse{}, err
	}

	return decoded, nil
}

func (p *Platform) doFetch(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, errors.New("status not 200")
	}

	return res, nil
}

func filterJob(job detailJob, opts search.Options) bool {
	for _, keyword := range opts.Keywords {
		if strings.Contains(job.Detail.Intro, keyword) {
			return true
		}
		if strings.Contains(job.Detail.Requirements, keyword) {
			return true
		}
		if strings.Contains(job.Detail.PreferredPoints, keyword) {
			return true
		}
		if strings.Contains(job.Detail.MainTasks, keyword) {
			return true
		}
		if strings.Contains(job.Detail.Position, keyword) {
			return true
		}
		if strings.Contains(job.Detail.Benefits, keyword) {
			return true
		}
		if strings.Contains(job.Detail.HireRounds, keyword) {
			return true
		}
	}

	return false
}

func makeJobDetailURL(id int) string {
	const base = "https://www.wanted.co.kr/api/chaos/jobs/v2"
	return strings.Join([]string{base, strconv.Itoa(id), "details"}, "/")
}

func makeJobDetailFrontendURL(id int) string {
	const base = "https://www.wanted.co.kr/wd"
	return strings.Join([]string{base, strconv.Itoa(id)}, "/")
}
