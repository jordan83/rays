package main

type CameraDef struct {
	Eye          Point3D
	LookAt       Point3D
	Up           Vector3D
	Roll         float32
	ExposureTime float32
}

// Computes an orthonormal basis for the camera, returns vectors, u, v, w (normalized)
func (c CameraDef) ComputeOrthonormalBasis() (Vector3D, Vector3D, Vector3D) {
	w := c.Eye.Subtract(c.LookAt).Normalize()
	u := c.Up.Cross(w).Normalize()
	v := w.Cross(u).Normalize()

	return u, v, w
}

type PinholeCamera struct {
	viewPlaneDistance float64
	zoomFactor        float64
	cameraDef         CameraDef
	u, v, w           Vector3D
}

func NewPinholeCamera(cameraDef CameraDef, viewPlaneDistance, zoomFactor float64) PinholeCamera {
	u, v, w := cameraDef.ComputeOrthonormalBasis()
	return PinholeCamera{
		viewPlaneDistance,
		zoomFactor,
		cameraDef,
		u,
		v,
		w,
	}
}

func (cam PinholeCamera) RenderScene(w *World, callback RenderCallback) {
	vp := w.GetViewPlane()
	ray := Ray{Origin: cam.cameraDef.Eye}
	pixelSize := float64(vp.PixelSize) / cam.zoomFactor
	pp := NewPoint3D()

	for r := 0; r < w.viewPlane.Vres; r++ {
		for c := 0; c < w.viewPlane.Hres; c++ {

			pixelColor := Black()
			for j := 0; j < vp.NumSamples; j++ {
				sample := w.GetSampler().SampleUnitSquare()

				pp.X = pixelSize * (float64(c) - 0.5*float64(vp.Hres) + sample.X)
				pp.Y = pixelSize * (float64(r) - 0.5*float64(vp.Vres) + sample.Y)
				ray.Direction = cam.getDirection(pp)

				pixelColor = pixelColor.Add(w.GetTracer().TraceRay(&ray))
			}

			callback(r, c, pixelColor.DivideBy(float64(vp.NumSamples)))
		}
	}
}

// Gets the direction of the point given the camera position
func (c PinholeCamera) getDirection(p Point3D) Vector3D {
	return c.u.ScalarMul(p.X).Add(c.v.ScalarMul(p.Y)).Subtract(c.w.ScalarMul(c.viewPlaneDistance))
}
