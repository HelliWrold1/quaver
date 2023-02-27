package ffmpeg

import (
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
				src:       "/home/helliwrold1/go/src/github.com/HelliWrold1/quaver/static/videos/1.mp4",
				dst:       "/home/helliwrold1/go/src/github.com/HelliWrold1/quaver/static/images/1.jpg",
				duration:  1 * time.Second,
				overwrite: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Get().Thumbnail(tt.args.src, tt.args.dst, tt.args.duration, tt.args.overwrite); (err != nil) != tt.wantErr {
				t.Errorf("Thumbnail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
