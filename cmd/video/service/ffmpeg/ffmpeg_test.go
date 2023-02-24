package ffmpeg

import (
	"os"
	"testing"
	"time"
)

func Test_Thumbnail(t *testing.T) {

	type args struct {
		src       string
		dst       string
		duration  time.Duration
		overwrite bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				//src:       "../../../../static/videos/1.mp4",
				//dst:       "../../../../static/images/1.jpg",
				src:       "../static/videos/1.mp4",
				dst:       "../static/images/1.jpg",
				duration:  1 * time.Second,
				overwrite: true,
			},
			wantErr: false,
		},
	}
	pwd, _ := os.Getwd()
	t.Log(pwd)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Get().Thumbnail(tt.args.src, tt.args.dst, tt.args.duration, tt.args.overwrite); (err != nil) != tt.wantErr {
				t.Errorf("Thumbnail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
