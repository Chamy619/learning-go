package solver

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type Processor struct {
	Solver MathSolver
}

type MathSolver interface {
	Resolve(ctx context.Context, expression string) (float64, error)
}

type RemoteSolver struct {
	MathServerURL string
	Client        *http.Client
}

func (rs RemoteSolver) Resolve(ctx context.Context, expression string) (float64, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rs.MathServerURL+"?expression="+url.QueryEscape(expression), nil)
	if err != nil {
		return 0, err
	}

	resp, err := rs.Client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil
	}
	if resp.StatusCode != http.StatusOK {
		return 0, errors.New(string(contents))
	}
	result, err := strconv.ParseFloat(string(contents), 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (p Processor) ProcessExpression(ctx context.Context, r io.Reader) (float64, error) {
	curExpression, err := readToNewLine(r)
	if err != nil {
		return 0, nil
	}

	if len(curExpression) == 0 {
		return 0, errors.New("no expression to read")
	}
	answer, err := p.Solver.Resolve(ctx, curExpression)
	return answer, err
}

func readToNewLine(r io.Reader) (string, error) {
	var out []byte
	b := make([]byte, 1)
	for {
		_, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				return string(out), nil
			}
		}
		if b[0] == '\n' {
			break
		}
		out = append(out, b[0])
	}
	return string(out), nil
}
