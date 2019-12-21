# Transform

[![GoDoc](https://godoc.org/github.com/bake/transform?status.svg)](https://godoc.org/github.com/bake/transform)

Rotate, Scale, Shear and Translate [`draw.Image`](https://godoc.org/image/draw#Image)s.

| Function | Output |
| --- | --- |
| Original | ![Original](assets/original.png) |
| `transform.Rotate(img, 45)` | ![Rotate](./assets/rotate-45.png) |
| `transform.Rotate(img, 90)` | ![Rotate](./assets/rotate-90.png) |
| `transform.Shear(img, 1.25, 0)`| ![Shear](assets/shear-1.25-0.png) |
| `transform.Translate(img, 150, 150)`| ![Shear](assets/translate-150-150.png) |
| `transform.Scale(img, 0.75, 0.75)` | ![Scale](./assets/scale-0.75-0.75.png) |
| `transform.Scale(img, 1.5, 1.5)`<br>`transform.Rotate(img, 45)` | ![Scale](./assets/scale-1.5-1.5-rotate-45.png) |
| `transform.MirrorX(img)` | ![MirrorX](./assets/mirrorx.png) |
| `transform.MirrorY(img)` | ![MirrorY](assets/mirrory.png) |
