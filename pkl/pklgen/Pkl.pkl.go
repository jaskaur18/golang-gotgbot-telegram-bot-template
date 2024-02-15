// Code generated from Pkl module `botConfig.pkl`. DO NOT EDIT.
package pklgen

import (
	"context"

	"github.com/apple/pkl-go/pkl"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/pkl/pklgen/environment"
)

type Pkl struct {
	Bot *BOT `pkl:"Bot"`

	DB *Database `pkl:"DB"`

	Logger *LOGGER `pkl:"Logger"`

	ENV environment.Environment `pkl:"ENV"`

	Misc *MISC `pkl:"Misc"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Pkl
func LoadFromPath(ctx context.Context, path string) (ret *Pkl, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Pkl
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Pkl, error) {
	var ret Pkl
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
