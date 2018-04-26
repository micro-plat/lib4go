package images

import (
	"testing"
)

func TestNewDraw(t *testing.T) {
	img := NewImage(100, 100)
	if img == nil {
		t.Error("test fail")
	}

	img = NewImage(0, 0)
	if img == nil {
		t.Error("test fail")
	}

	img = NewImage(-100, 100)
	if img == nil {
		t.Error("test fail")
	}

	img = NewImage(-100, -100)
	if img == nil {
		t.Error("test fail")
	}

	img = NewImage(100, -100)
	if img == nil {
		t.Error("test fail")
	}
}

func TestNewImageFromFile(t *testing.T) {
	path := "/home/champly/picture.png"
	img, err := NewImageFromFile(100, 100, path)
	if err != nil {
		t.Errorf("test fail %v", err)
	}
	if img == nil {
		t.Error("test fail")
	}

	path = "/home/champly/picture.jpg"
	img, err = NewImageFromFile(-100, 100, path)
	if err != nil {
		t.Errorf("test fail %v", err)
	}
	if img == nil {
		t.Error("test fail")
	}

	path = "/home/champly/picture.jpg"
	img, err = NewImageFromFile(-100, -100, path)
	if err != nil {
		t.Errorf("test fail %v", err)
	}
	if img == nil {
		t.Error("test fail")
	}

	path = "/home/champly/picture.jpg"
	img, err = NewImageFromFile(100, -100, path)
	if err != nil {
		t.Errorf("test fail %v", err)
	}
	if img == nil {
		t.Error("test fail")
	}

	path = "/home/champly/err_picture.jpg"
	img, err = NewImageFromFile(100, 100, path)
	if err == nil {
		t.Error("test fail")
	}

	path = "/home/champly/picture.gif"
	img, err = NewImageFromFile(100, 100, path)
	if err == nil {
		t.Error("test fail")
	}
}

func TestDrawFont(t *testing.T) {
	path := "/home/champly/picture.jpg"
	img, err := NewImageFromFile(1920, 1080, path)
	if err != nil {
		t.Errorf("test fail %v", err)
	}
	if img == nil {
		t.Error("test fail")
	}

	fontPath := "/usr/share/fonts/msyhbd.ttf"
	text := "Hello World"
	col := "155"
	fontSize := 16.0
	img.DrawFont(fontPath, text, col, fontSize, 100, 300)
	err = img.Save("/home/champly/picture_test.png")
	if err != nil {
		t.Errorf("test fail %v", err)
	}

	fontPath = "/usr/share/fonts/msyhbdsdfadf.ttf"
	text = "Hello World"
	col = "155"
	fontSize = 16.0
	img.DrawFont(fontPath, text, col, fontSize, 100, 300)
	err = img.Save("/home/champly/picture_test.png")
	if err != nil {
		t.Errorf("test fail %v", err)
	}
}

func TestDrawImage(t *testing.T) {
	path := "/home/champly/picture_test.png"
	img, err := NewImageFromFile(1920, 1080, path)
	if err != nil {
		t.Errorf("test fail %v", err)
	}
	if img == nil {
		t.Error("test fail")
	}

	err = img.Save("/home/champly/picture_test.png")
	if err != nil {
		t.Errorf("test fail %v", err)
	}

	// 路径不正确保存文件
	err = img.Save("/home/champly/picture/picture_test.png")
	if err == nil {
		t.Error("test fail")
	}
}
