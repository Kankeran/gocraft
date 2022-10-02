package noise

// This is the new and improved, C(2) continuous interpolant
func FADE(t float64) float64 {
	return t * t * t * (t*(t*6-15) + 10)
}

func FASTFLOOR(x float64) int {
	if float64(int(x)) < x {
		return int(x)
	}
	return int(x) - 1
}

func LERP(t, a, b float64) float64 {
	return a + t*(b-a)
}

var perm = []int{151, 160, 137, 91, 90, 15,
	131, 13, 201, 95, 96, 53, 194, 233, 7, 225, 140, 36, 103, 30, 69, 142, 8, 99, 37, 240, 21, 10, 23,
	190, 6, 148, 247, 120, 234, 75, 0, 26, 197, 62, 94, 252, 219, 203, 117, 35, 11, 32, 57, 177, 33,
	88, 237, 149, 56, 87, 174, 20, 125, 136, 171, 168, 68, 175, 74, 165, 71, 134, 139, 48, 27, 166,
	77, 146, 158, 231, 83, 111, 229, 122, 60, 211, 133, 230, 220, 105, 92, 41, 55, 46, 245, 40, 244,
	102, 143, 54, 65, 25, 63, 161, 1, 216, 80, 73, 209, 76, 132, 187, 208, 89, 18, 169, 200, 196,
	135, 130, 116, 188, 159, 86, 164, 100, 109, 198, 173, 186, 3, 64, 52, 217, 226, 250, 124, 123,
	5, 202, 38, 147, 118, 126, 255, 82, 85, 212, 207, 206, 59, 227, 47, 16, 58, 17, 182, 189, 28, 42,
	223, 183, 170, 213, 119, 248, 152, 2, 44, 154, 163, 70, 221, 153, 101, 155, 167, 43, 172, 9,
	129, 22, 39, 253, 19, 98, 108, 110, 79, 113, 224, 232, 178, 185, 112, 104, 218, 246, 97, 228,
	251, 34, 242, 193, 238, 210, 144, 12, 191, 179, 162, 241, 81, 51, 145, 235, 249, 14, 239, 107,
	49, 192, 214, 31, 181, 199, 106, 157, 184, 84, 204, 176, 115, 121, 50, 45, 127, 4, 150, 254,
	138, 236, 205, 93, 222, 114, 67, 29, 24, 72, 243, 141, 128, 195, 78, 66, 215, 61, 156, 180,
	151, 160, 137, 91, 90, 15,
	131, 13, 201, 95, 96, 53, 194, 233, 7, 225, 140, 36, 103, 30, 69, 142, 8, 99, 37, 240, 21, 10, 23,
	190, 6, 148, 247, 120, 234, 75, 0, 26, 197, 62, 94, 252, 219, 203, 117, 35, 11, 32, 57, 177, 33,
	88, 237, 149, 56, 87, 174, 20, 125, 136, 171, 168, 68, 175, 74, 165, 71, 134, 139, 48, 27, 166,
	77, 146, 158, 231, 83, 111, 229, 122, 60, 211, 133, 230, 220, 105, 92, 41, 55, 46, 245, 40, 244,
	102, 143, 54, 65, 25, 63, 161, 1, 216, 80, 73, 209, 76, 132, 187, 208, 89, 18, 169, 200, 196,
	135, 130, 116, 188, 159, 86, 164, 100, 109, 198, 173, 186, 3, 64, 52, 217, 226, 250, 124, 123,
	5, 202, 38, 147, 118, 126, 255, 82, 85, 212, 207, 206, 59, 227, 47, 16, 58, 17, 182, 189, 28, 42,
	223, 183, 170, 213, 119, 248, 152, 2, 44, 154, 163, 70, 221, 153, 101, 155, 167, 43, 172, 9,
	129, 22, 39, 253, 19, 98, 108, 110, 79, 113, 224, 232, 178, 185, 112, 104, 218, 246, 97, 228,
	251, 34, 242, 193, 238, 210, 144, 12, 191, 179, 162, 241, 81, 51, 145, 235, 249, 14, 239, 107,
	49, 192, 214, 31, 181, 199, 106, 157, 184, 84, 204, 176, 115, 121, 50, 45, 127, 4, 150, 254,
	138, 236, 205, 93, 222, 114, 67, 29, 24, 72, 243, 141, 128, 195, 78, 66, 215, 61, 156, 180,
}

func grad1(hash int, x float64) float64 {
	h := hash & 15
	grad := 1.0 + float64(h&7)
	if h&8 > 0 {
		grad = -grad
	}
	return grad * x
}

func grad2(hash int, x, y float64) float64 {
	h := hash & 7
	u, v := y, x
	if h < 4 {
		u, v = x, y
	}
	if h&1 > 0 {
		u = -u
	}
	if h&2 > 0 {
		v = -v
	}
	return u + 2*v
}

func grad3(hash int, x, y, z float64) float64 {
	h := hash & 15
	u, v := y, z
	if h < 8 {
		u = x
	}
	if h < 4 {
		v = y
	} else if h == 12 || h == 14 {
		v = x
	}

	if h&1 > 0 {
		u = -u
	}
	if h&2 > 0 {
		v = -v
	}
	return u + v
}

func grad4(hash int, x, y, z, t float64) float64 {
	h := hash & 31
	u, v, w := y, z, t
	if h < 24 {
		u = x
	}
	if h < 16 {
		v = y
	}
	if h < 8 {
		w = z
	}
	if h&1 > 0 {
		u = -u
	}
	if h&2 > 0 {
		v = -v
	}
	if h&4 > 0 {
		w = -w
	}

	return u + v + w
}

//---------------------------------------------------------------------
/** 1D float Perlin noise, SL "noise()"
 */
func Noise1(x float64) float64 {
	ix0 := FASTFLOOR(x)     // Integer part of x
	fx0 := x - float64(ix0) // Fractional part of x
	fx1 := fx0 - 1.0
	ix1 := (ix0 + 1) & 0xff
	ix0 = ix0 & 0xff // Wrap to 0..255

	s := FADE(fx0)

	n0 := grad1(perm[ix0], fx0)
	n1 := grad1(perm[ix1], fx1)
	return 0.188 * (LERP(s, n0, n1))
}

//---------------------------------------------------------------------
/** 1D float Perlin periodic noise, SL "pnoise()"
 */
func Pnoise1(x float64, px int) float64 {
	ix0 := FASTFLOOR(x)     // Integer part of x
	fx0 := x - float64(ix0) // Fractional part of x
	fx1 := fx0 - 1.0
	ix1 := ((ix0 + 1) % px) & 0xff // Wrap to 0..px-1 *and* wrap to 0..255
	ix0 = (ix0 % px) & 0xff        // (because px might be greater than 256)

	s := FADE(fx0)

	n0 := grad1(perm[ix0], fx0)
	n1 := grad1(perm[ix1], fx1)
	return 0.188 * (LERP(s, n0, n1))
}

//---------------------------------------------------------------------
/** 2D float Perlin noise.
 */
func Noise2(x, y float64) float64 {
	ix0 := FASTFLOOR(x)     // Integer part of x
	iy0 := FASTFLOOR(y)     // Integer part of y
	fx0 := x - float64(ix0) // Fractional part of x
	fy0 := y - float64(iy0) // Fractional part of y
	fx1 := fx0 - 1.0
	fy1 := fy0 - 1.0
	ix1 := (ix0 + 1) & 0xff // Wrap to 0..255
	iy1 := (iy0 + 1) & 0xff
	ix0 = ix0 & 0xff
	iy0 = iy0 & 0xff

	t := FADE(fy0)
	s := FADE(fx0)

	nx0 := grad2(perm[ix0+perm[iy0]], fx0, fy0)
	nx1 := grad2(perm[ix0+perm[iy1]], fx0, fy1)
	n0 := LERP(t, nx0, nx1)

	nx0 = grad2(perm[ix1+perm[iy0]], fx1, fy0)
	nx1 = grad2(perm[ix1+perm[iy1]], fx1, fy1)
	n1 := LERP(t, nx0, nx1)

	return 0.507 * (LERP(s, n0, n1))
}

//---------------------------------------------------------------------
/** 2D float Perlin periodic noise.
 */
func Pnoise2(x, y float64, px, py int) float64 {
	ix0 := FASTFLOOR(x)     // Integer part of x
	iy0 := FASTFLOOR(y)     // Integer part of y
	fx0 := x - float64(ix0) // Fractional part of x
	fy0 := y - float64(iy0) // Fractional part of y
	fx1 := fx0 - 1.0
	fy1 := fy0 - 1.0
	ix1 := ((ix0 + 1) % px) & 0xff // Wrap to 0..px-1 and wrap to 0..255
	iy1 := ((iy0 + 1) % py) & 0xff // Wrap to 0..py-1 and wrap to 0..255
	ix0 = (ix0 % px) & 0xff
	iy0 = (iy0 % py) & 0xff

	t := FADE(fy0)
	s := FADE(fx0)

	nx0 := grad2(perm[ix0+perm[iy0]], fx0, fy0)
	nx1 := grad2(perm[ix0+perm[iy1]], fx0, fy1)
	n0 := LERP(t, nx0, nx1)

	nx0 = grad2(perm[ix1+perm[iy0]], fx1, fy0)
	nx1 = grad2(perm[ix1+perm[iy1]], fx1, fy1)
	n1 := LERP(t, nx0, nx1)

	return 0.507 * (LERP(s, n0, n1))
}

//---------------------------------------------------------------------
/** 3D float Perlin noise.
 */
func Noise3(x, y, z float64) float64 {
	ix0 := FASTFLOOR(x)     // Integer part of x
	iy0 := FASTFLOOR(y)     // Integer part of y
	iz0 := FASTFLOOR(z)     // Integer part of z
	fx0 := x - float64(ix0) // Fractional part of x
	fy0 := y - float64(iy0) // Fractional part of y
	fz0 := z - float64(iz0) // Fractional part of z
	fx1 := fx0 - 1.0
	fy1 := fy0 - 1.0
	fz1 := fz0 - 1.0
	ix1 := (ix0 + 1) & 0xff // Wrap to 0..255
	iy1 := (iy0 + 1) & 0xff
	iz1 := (iz0 + 1) & 0xff
	ix0 = ix0 & 0xff
	iy0 = iy0 & 0xff
	iz0 = iz0 & 0xff

	r := FADE(fz0)
	t := FADE(fy0)
	s := FADE(fx0)

	nxy0 := grad3(perm[ix0+perm[iy0+perm[iz0]]], fx0, fy0, fz0)
	nxy1 := grad3(perm[ix0+perm[iy0+perm[iz1]]], fx0, fy0, fz1)
	nx0 := LERP(r, nxy0, nxy1)

	nxy0 = grad3(perm[ix0+perm[iy1+perm[iz0]]], fx0, fy1, fz0)
	nxy1 = grad3(perm[ix0+perm[iy1+perm[iz1]]], fx0, fy1, fz1)
	nx1 := LERP(r, nxy0, nxy1)

	n0 := LERP(t, nx0, nx1)

	nxy0 = grad3(perm[ix1+perm[iy0+perm[iz0]]], fx1, fy0, fz0)
	nxy1 = grad3(perm[ix1+perm[iy0+perm[iz1]]], fx1, fy0, fz1)
	nx0 = LERP(r, nxy0, nxy1)

	nxy0 = grad3(perm[ix1+perm[iy1+perm[iz0]]], fx1, fy1, fz0)
	nxy1 = grad3(perm[ix1+perm[iy1+perm[iz1]]], fx1, fy1, fz1)
	nx1 = LERP(r, nxy0, nxy1)

	n1 := LERP(t, nx0, nx1)

	return 0.936 * (LERP(s, n0, n1))
}

//---------------------------------------------------------------------
/** 3D float Perlin periodic noise.
 */
func Pnoise3(x, y, z float64, px, py, pz int) float64 {
	ix0 := FASTFLOOR(x)     // Integer part of x
	iy0 := FASTFLOOR(y)     // Integer part of y
	iz0 := FASTFLOOR(z)     // Integer part of z
	fx0 := x - float64(ix0) // Fractional part of x
	fy0 := y - float64(iy0) // Fractional part of y
	fz0 := z - float64(iz0) // Fractional part of z
	fx1 := fx0 - 1.0
	fy1 := fy0 - 1.0
	fz1 := fz0 - 1.0
	ix1 := ((ix0 + 1) % px) & 0xff // Wrap to 0..px-1 and wrap to 0..255
	iy1 := ((iy0 + 1) % py) & 0xff // Wrap to 0..py-1 and wrap to 0..255
	iz1 := ((iz0 + 1) % pz) & 0xff // Wrap to 0..pz-1 and wrap to 0..255
	ix0 = (ix0 % px) & 0xff
	iy0 = (iy0 % py) & 0xff
	iz0 = (iz0 % pz) & 0xff

	r := FADE(fz0)
	t := FADE(fy0)
	s := FADE(fx0)

	nxy0 := grad3(perm[ix0+perm[iy0+perm[iz0]]], fx0, fy0, fz0)
	nxy1 := grad3(perm[ix0+perm[iy0+perm[iz1]]], fx0, fy0, fz1)
	nx0 := LERP(r, nxy0, nxy1)

	nxy0 = grad3(perm[ix0+perm[iy1+perm[iz0]]], fx0, fy1, fz0)
	nxy1 = grad3(perm[ix0+perm[iy1+perm[iz1]]], fx0, fy1, fz1)
	nx1 := LERP(r, nxy0, nxy1)

	n0 := LERP(t, nx0, nx1)

	nxy0 = grad3(perm[ix1+perm[iy0+perm[iz0]]], fx1, fy0, fz0)
	nxy1 = grad3(perm[ix1+perm[iy0+perm[iz1]]], fx1, fy0, fz1)
	nx0 = LERP(r, nxy0, nxy1)

	nxy0 = grad3(perm[ix1+perm[iy1+perm[iz0]]], fx1, fy1, fz0)
	nxy1 = grad3(perm[ix1+perm[iy1+perm[iz1]]], fx1, fy1, fz1)
	nx1 = LERP(r, nxy0, nxy1)

	n1 := LERP(t, nx0, nx1)

	return 0.936 * (LERP(s, n0, n1))
}

//---------------------------------------------------------------------
/** 4D float Perlin noise.
 */

func Noise4(x, y, z, w float64) float64 {
	ix0 := FASTFLOOR(x)     // Integer part of x
	iy0 := FASTFLOOR(y)     // Integer part of y
	iz0 := FASTFLOOR(z)     // Integer part of y
	iw0 := FASTFLOOR(w)     // Integer part of w
	fx0 := x - float64(ix0) // Fractional part of x
	fy0 := y - float64(iy0) // Fractional part of y
	fz0 := z - float64(iz0) // Fractional part of z
	fw0 := w - float64(iw0) // Fractional part of w
	fx1 := fx0 - 1.0
	fy1 := fy0 - 1.0
	fz1 := fz0 - 1.0
	fw1 := fw0 - 1.0
	ix1 := (ix0 + 1) & 0xff // Wrap to 0..255
	iy1 := (iy0 + 1) & 0xff
	iz1 := (iz0 + 1) & 0xff
	iw1 := (iw0 + 1) & 0xff
	ix0 = ix0 & 0xff
	iy0 = iy0 & 0xff
	iz0 = iz0 & 0xff
	iw0 = iw0 & 0xff

	q := FADE(fw0)
	r := FADE(fz0)
	t := FADE(fy0)
	s := FADE(fx0)

	nxyz0 := grad4(perm[ix0+perm[iy0+perm[iz0+perm[iw0]]]], fx0, fy0, fz0, fw0)
	nxyz1 := grad4(perm[ix0+perm[iy0+perm[iz0+perm[iw1]]]], fx0, fy0, fz0, fw1)
	nxy0 := LERP(q, nxyz0, nxyz1)

	nxyz0 = grad4(perm[ix0+perm[iy0+perm[iz1+perm[iw0]]]], fx0, fy0, fz1, fw0)
	nxyz1 = grad4(perm[ix0+perm[iy0+perm[iz1+perm[iw1]]]], fx0, fy0, fz1, fw1)
	nxy1 := LERP(q, nxyz0, nxyz1)

	nx0 := LERP(r, nxy0, nxy1)

	nxyz0 = grad4(perm[ix0+perm[iy1+perm[iz0+perm[iw0]]]], fx0, fy1, fz0, fw0)
	nxyz1 = grad4(perm[ix0+perm[iy1+perm[iz0+perm[iw1]]]], fx0, fy1, fz0, fw1)
	nxy0 = LERP(q, nxyz0, nxyz1)

	nxyz0 = grad4(perm[ix0+perm[iy1+perm[iz1+perm[iw0]]]], fx0, fy1, fz1, fw0)
	nxyz1 = grad4(perm[ix0+perm[iy1+perm[iz1+perm[iw1]]]], fx0, fy1, fz1, fw1)
	nxy1 = LERP(q, nxyz0, nxyz1)

	nx1 := LERP(r, nxy0, nxy1)

	n0 := LERP(t, nx0, nx1)

	nxyz0 = grad4(perm[ix1+perm[iy0+perm[iz0+perm[iw0]]]], fx1, fy0, fz0, fw0)
	nxyz1 = grad4(perm[ix1+perm[iy0+perm[iz0+perm[iw1]]]], fx1, fy0, fz0, fw1)
	nxy0 = LERP(q, nxyz0, nxyz1)

	nxyz0 = grad4(perm[ix1+perm[iy0+perm[iz1+perm[iw0]]]], fx1, fy0, fz1, fw0)
	nxyz1 = grad4(perm[ix1+perm[iy0+perm[iz1+perm[iw1]]]], fx1, fy0, fz1, fw1)
	nxy1 = LERP(q, nxyz0, nxyz1)

	nx0 = LERP(r, nxy0, nxy1)

	nxyz0 = grad4(perm[ix1+perm[iy1+perm[iz0+perm[iw0]]]], fx1, fy1, fz0, fw0)
	nxyz1 = grad4(perm[ix1+perm[iy1+perm[iz0+perm[iw1]]]], fx1, fy1, fz0, fw1)
	nxy0 = LERP(q, nxyz0, nxyz1)

	nxyz0 = grad4(perm[ix1+perm[iy1+perm[iz1+perm[iw0]]]], fx1, fy1, fz1, fw0)
	nxyz1 = grad4(perm[ix1+perm[iy1+perm[iz1+perm[iw1]]]], fx1, fy1, fz1, fw1)
	nxy1 = LERP(q, nxyz0, nxyz1)

	nx1 = LERP(r, nxy0, nxy1)

	n1 := LERP(t, nx0, nx1)

	return 0.87 * (LERP(s, n0, n1))
}

//---------------------------------------------------------------------
/** 4D float Perlin periodic noise.
 */

func Pnoise4(x, y, z, w float64, px, py, pz, pw int) float64 {
	ix0 := FASTFLOOR(x)     // Integer part of x
	iy0 := FASTFLOOR(y)     // Integer part of y
	iz0 := FASTFLOOR(z)     // Integer part of y
	iw0 := FASTFLOOR(w)     // Integer part of w
	fx0 := x - float64(ix0) // Fractional part of x
	fy0 := y - float64(iy0) // Fractional part of y
	fz0 := z - float64(iz0) // Fractional part of z
	fw0 := w - float64(iw0) // Fractional part of w
	fx1 := fx0 - 1.0
	fy1 := fy0 - 1.0
	fz1 := fz0 - 1.0
	fw1 := fw0 - 1.0
	ix1 := ((ix0 + 1) % px) & 0xff // Wrap to 0..px-1 and wrap to 0..255
	iy1 := ((iy0 + 1) % py) & 0xff // Wrap to 0..py-1 and wrap to 0..255
	iz1 := ((iz0 + 1) % pz) & 0xff // Wrap to 0..pz-1 and wrap to 0..255
	iw1 := ((iw0 + 1) % pw) & 0xff // Wrap to 0..pw-1 and wrap to 0..255
	ix0 = (ix0 % px) & 0xff
	iy0 = (iy0 % py) & 0xff
	iz0 = (iz0 % pz) & 0xff
	iw0 = (iw0 % pw) & 0xff

	q := FADE(fw0)
	r := FADE(fz0)
	t := FADE(fy0)
	s := FADE(fx0)

	nxyz0 := grad4(perm[ix0+perm[iy0+perm[iz0+perm[iw0]]]], fx0, fy0, fz0, fw0)
	nxyz1 := grad4(perm[ix0+perm[iy0+perm[iz0+perm[iw1]]]], fx0, fy0, fz0, fw1)
	nxy0 := LERP(q, nxyz0, nxyz1)

	nxyz0 = grad4(perm[ix0+perm[iy0+perm[iz1+perm[iw0]]]], fx0, fy0, fz1, fw0)
	nxyz1 = grad4(perm[ix0+perm[iy0+perm[iz1+perm[iw1]]]], fx0, fy0, fz1, fw1)
	nxy1 := LERP(q, nxyz0, nxyz1)

	nx0 := LERP(r, nxy0, nxy1)

	nxyz0 = grad4(perm[ix0+perm[iy1+perm[iz0+perm[iw0]]]], fx0, fy1, fz0, fw0)
	nxyz1 = grad4(perm[ix0+perm[iy1+perm[iz0+perm[iw1]]]], fx0, fy1, fz0, fw1)
	nxy0 = LERP(q, nxyz0, nxyz1)

	nxyz0 = grad4(perm[ix0+perm[iy1+perm[iz1+perm[iw0]]]], fx0, fy1, fz1, fw0)
	nxyz1 = grad4(perm[ix0+perm[iy1+perm[iz1+perm[iw1]]]], fx0, fy1, fz1, fw1)
	nxy1 = LERP(q, nxyz0, nxyz1)

	nx1 := LERP(r, nxy0, nxy1)

	n0 := LERP(t, nx0, nx1)

	nxyz0 = grad4(perm[ix1+perm[iy0+perm[iz0+perm[iw0]]]], fx1, fy0, fz0, fw0)
	nxyz1 = grad4(perm[ix1+perm[iy0+perm[iz0+perm[iw1]]]], fx1, fy0, fz0, fw1)
	nxy0 = LERP(q, nxyz0, nxyz1)

	nxyz0 = grad4(perm[ix1+perm[iy0+perm[iz1+perm[iw0]]]], fx1, fy0, fz1, fw0)
	nxyz1 = grad4(perm[ix1+perm[iy0+perm[iz1+perm[iw1]]]], fx1, fy0, fz1, fw1)
	nxy1 = LERP(q, nxyz0, nxyz1)

	nx0 = LERP(r, nxy0, nxy1)

	nxyz0 = grad4(perm[ix1+perm[iy1+perm[iz0+perm[iw0]]]], fx1, fy1, fz0, fw0)
	nxyz1 = grad4(perm[ix1+perm[iy1+perm[iz0+perm[iw1]]]], fx1, fy1, fz0, fw1)
	nxy0 = LERP(q, nxyz0, nxyz1)

	nxyz0 = grad4(perm[ix1+perm[iy1+perm[iz1+perm[iw0]]]], fx1, fy1, fz1, fw0)
	nxyz1 = grad4(perm[ix1+perm[iy1+perm[iz1+perm[iw1]]]], fx1, fy1, fz1, fw1)
	nxy1 = LERP(q, nxyz0, nxyz1)

	nx1 = LERP(r, nxy0, nxy1)

	n1 := LERP(t, nx0, nx1)

	return 0.87 * (LERP(s, n0, n1))
}
