package assets

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	textv2 "github.com/hajimehoshi/ebiten/v2/text/v2"
)

const debugFontPath = "assets/fonts/msyh.ttc"

var (
	debugFontSourceOnce sync.Once
	debugFontSource     *textv2.GoTextFaceSource
	debugFontSourceErr  error
)

// DebugTextFace 返回调试面板使用的字体
func DebugTextFace(size float64) (*textv2.GoTextFace, error) {
	source, err := debugFontFaceSource()
	if err != nil {
		return nil, err
	}

	return &textv2.GoTextFace{
		Source: source,
		Size:   size,
	}, nil
}

// debugFontFaceSource 懒加载调试面板使用的字体源
func debugFontFaceSource() (*textv2.GoTextFaceSource, error) {
	debugFontSourceOnce.Do(func() {
		data, err := os.ReadFile(filepath.Clean(debugFontPath))
		if err != nil {
			debugFontSourceErr = err
			return
		}

		sources, err := textv2.NewGoTextFaceSourcesFromCollection(bytes.NewReader(data))
		if err == nil && len(sources) > 0 {
			debugFontSource = sources[0]
			return
		}

		source, singleErr := textv2.NewGoTextFaceSource(bytes.NewReader(data))
		if singleErr != nil {
			if err != nil {
				debugFontSourceErr = fmt.Errorf("parse font collection: %w; parse single font: %v", err, singleErr)
				return
			}
			debugFontSourceErr = singleErr
			return
		}

		debugFontSource = source
	})

	return debugFontSource, debugFontSourceErr
}
