package sdl

// #cgo pkg-config: SDL_gfx
// #include <SDL_rotozoom.h>
import "C"

func (s *Surface) Zoom(zoomX, zoomY float64, smooth bool) *Surface {
	cSmooth := C.int(0)
	if smooth {
		cSmooth = C.int(1)
	}
	return wrap(C.zoomSurface(s.cSurface, C.double(zoomX), C.double(zoomY), cSmooth))
}
