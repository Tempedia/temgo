package files

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/api/middleware"
	"gitlab.com/wiky.lyu/temgo/service/files"
)

type UploadRequest struct {
	From string `json:"from" form:"from" query:"from"`
}

func Upload(c echo.Context) error {
	ctx := c.(*middleware.Context)

	req := UploadRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.BadRequest()
	}

	file, err := ctx.FormFile(`file`)
	if err != nil {
		return ctx.BadRequest()
	}
	reader, err := file.Open()
	if err != nil {
		return ctx.BadRequest()
	}
	defer reader.Close()

	tmpfile, err := os.CreateTemp(files.Folder(), "temgo_upload_")
	if err != nil {
		return err
	}
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	m := md5.New()
	buf := make([]byte, 4096)
	for {
		if n, err := reader.Read(buf); err != nil {
			if err == io.EOF {
				break
			}
			return err
		} else if n <= 0 {
			break
		} else {
			if _, err := tmpfile.Write(buf[:n]); err != nil {
				return err
			}
			if _, err := m.Write(buf[:n]); err != nil {
				return err
			}
		}
	}
	filename := fmt.Sprintf("%x", m.Sum(nil))

	filepath := path.Join(files.Folder(), filename)
	if err := os.Rename(tmpfile.Name(), filepath); err != nil {
		return err
	}

	fileid := path.Base(filepath)
	if req.From == "tinymce" {
		/* https://www.tiny.cloud/docs/general-configuration-guide/upload-images/ */
		return ctx.JSON(200, map[string]string{
			"location": fileid,
		})
	}
	return ctx.Success(map[string]interface{}{
		"fileid": fileid,
	})
}
