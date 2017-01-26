package cloud

import (
	"github.com/ingrammicro/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGenericImageServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewGenericImageService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetGenericImageList(t *testing.T) {
	genericImagesIn := testdata.GetGenericImageData()
	GetGenericImageListMocked(t, genericImagesIn)
	GetGenericImageListFailErrMocked(t, genericImagesIn)
	GetGenericImageListFailStatusMocked(t, genericImagesIn)
	GetGenericImageListFailJSONMocked(t, genericImagesIn)
}
