// Code generated from Pkl module `MotrConfig`. DO NOT EDIT.
package motrconfig

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type MotrConfig struct {
	App *AppConfig `pkl:"app"`

	Bitcoin *BitcoinConfig `pkl:"bitcoin"`

	Cloudflare *CloudflareConfig `pkl:"cloudflare"`

	Ethereum *EthereumConfig `pkl:"ethereum"`

	Ipfs *IpfsConfig `pkl:"ipfs"`

	Sonr *SonrConfig `pkl:"sonr"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a MotrConfig
func LoadFromPath(ctx context.Context, path string) (ret *MotrConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a MotrConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*MotrConfig, error) {
	var ret MotrConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
