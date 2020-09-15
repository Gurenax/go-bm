package benchmark

import (
	"testing"
	"time"
)

var noTime time.Time

func TestBM_Start(t *testing.T) {
	type fields struct {
		startTime time.Time
		endTime   time.Time
		duration  time.Duration
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Executes successfully",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BM{
				startTime: tt.fields.startTime,
				endTime:   tt.fields.endTime,
				duration:  tt.fields.duration,
			}
			b.Start()
			if b.startTime == noTime {
				t.Errorf("Expected startTime to be initialised, but got %v", b.startTime)
			}
		})
	}
}

func TestBM_End(t *testing.T) {
	givenStartTime := time.Now()

	type fields struct {
		startTime time.Time
		endTime   time.Time
		duration  time.Duration
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Executes successfully",
			fields: fields{
				startTime: givenStartTime,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BM{
				startTime: tt.fields.startTime,
				endTime:   tt.fields.endTime,
				duration:  tt.fields.duration,
			}
			b.End()
			if b.startTime != givenStartTime {
				t.Errorf("Expected startTime to be initialised as %v, but got %v", givenStartTime, b.startTime)
			}
			if b.endTime == noTime {
				t.Errorf("Expected endTime to be initialised, but got %v", b.startTime)
			}
			if b.duration == 0 {
				t.Errorf("Expected duration to be initialised, but got %v", b.duration)
			}
		})
	}
}
