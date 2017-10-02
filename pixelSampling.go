package main

import (
	"math"
	"math/rand"
)

type Sampler struct {
	jump            int
	count           int
	numSamples      int
	numSets         int
	samples         []*Point3D
	shuffledIndexes []int
}

func NewSampler(numSamples int) *Sampler {
	numSets := 83
	var samples []*Point3D
	if numSamples > 1 {
		samples = multiJitteredSampler(numSets, numSamples)
	} else {
		samples = regularSampler(numSets, numSamples)
	}

	return &Sampler{0, 0, numSamples, numSets, samples, getShuffledIndexes(numSets, numSamples)}
}

func (s *Sampler) SampleUnitSquare() Point3D {
	if s.count%s.numSamples == 0 {
		s.jump = (rand.Int() % s.numSets) * s.numSamples
	}

	s.count += 1
	return *s.samples[s.jump+s.shuffledIndexes[s.jump+s.count%s.numSamples]]
}

func getShuffledIndexes(numSets, numSamples int) []int {
	shuffledIndexes := make([]int, 0, numSamples*numSets)

	for p := 0; p < numSets; p++ {
		indexes := rand.Perm(numSamples)

		for j := 0; j < numSamples; j++ {
			shuffledIndexes = append(shuffledIndexes, indexes[j])
		}
	}

	return shuffledIndexes
}

func regularSampler(numSets, numSamples int) []*Point3D {
	n := int(math.Sqrt(float64(numSamples)))
	points := make([]*Point3D, 0, numSamples*numSets)

	for p := 0; p < numSets; p++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				p := Point3D{(float64(k) + 0.5) / float64(n), (float64(j) + 0.5) / float64(n), 0}
				points = append(points, &p)
			}
		}
	}
	return points
}

func multiJitteredSampler(numSets, numSamples int) []*Point3D {
	n := int(math.Sqrt(float64(numSamples)))
	subcellWidth := 1.0 / float64(numSamples)

	points := make([]*Point3D, numSamples*numSets)
	for j := 0; j < numSamples*numSets; j++ {
		point := NewPoint3D()
		points[j] = &point
	}

	// Distribute points in initial pattern
	for p := 0; p < numSets; p++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				points[i*n+j+p*numSamples].X = float64(i*n+j)*subcellWidth + randFloat(0, subcellWidth)
				points[i*n+j+p*numSamples].Y = float64(j*n+i)*subcellWidth + randFloat(0, subcellWidth)
			}
		}
	}

	// shuffle x coordinates
	for p := 0; p < numSets; p++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				k := randInt(j, n-1)
				t := points[i*n+j+p*numSamples].X
				points[i*n+j+p*numSamples].X = points[i*n+k+p*numSamples].X
				points[i*n+k+p*numSamples].X = t
			}
		}
	}

	// shuffle y coordinates
	for p := 0; p < numSets; p++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				k := randInt(j, n-1)
				t := points[j*n+i+p*numSamples].Y
				points[j*n+i+p*numSamples].Y = points[k*n+i+p*numSamples].Y
				points[k*n+i+p*numSamples].Y = t
			}
		}
	}

	return points
}

func randFloat(low int, high float64) float64 {
	return rand.Float64()*(high-float64(low)) + float64(low)
}

func randInt(low, high int) int {
	return int(randFloat(0, float64(high-low+1))) + low
}
