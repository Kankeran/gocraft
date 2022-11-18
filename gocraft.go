package main

import (
	"fmt"
	"gocraft/control"
	"gocraft/graphics"
	"gocraft/world"
	"runtime"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

var drawMode uint32 = gl.FILL
var keyPressed = false
var cursorHasDisabled = false
var keyPressed2 = false
var keyPressed3 = false
var movementModel = 0

const windowWidth = 800
const windowHeight = 600

var fov float32 = 60.0
var drawMinDistance float32 = 0.1
var drawMaxDistance float32 = 1000.0

// camera
var speed float32 = 8
var camera *control.Camera

var shader *graphics.ShaderProgram

// var basicShader *graphics.ShaderProgram

var deltaTime, lastFrame float64

var thisWorld *world.World

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 5)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	// glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "gocraft", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	window.SetCursorPosCallback(cursorCallback)
	window.SetFramebufferSizeCallback(framebufferSizeCallback)
	window.SetMouseButtonCallback(mouseButtonCallback)

	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	var nrAttributes int32
	gl.GetIntegerv(gl.MAX_VERTEX_ATTRIBS, &nrAttributes)
	fmt.Println("Maximum nr of vertex attributes supported:", nrAttributes)

	shader = graphics.ProvideShaderProgramBasicTexture()
	shader.Use()
	shader.SetUniformMat4("projection\x00", mgl32.Perspective(mgl32.DegToRad(fov), float32(windowWidth)/float32(windowHeight), drawMinDistance, drawMaxDistance))
	shader.SetUniformTexture2D("tex\x00", graphics.ProvideTextureAtlas(), 0)

	thisWorld = world.ProvideWorld()
	thisWorld.Initialize()
	camera = control.ProvideCamera()

	// basicShader = container.GetService("ShaderProgram.basic").(*graphics.ShaderProgram)
	// basicShader.Use()
	// basicShader.SetUniformMat4("projection\x00", mgl32.Perspective(mgl32.DegToRad(fov), float32(windowWidth)/float32(windowHeight), drawMinDistance, drawMaxDistance))
	// var line = container.GetService("Line").(*graphics.Line)

	gl.Enable(gl.DEPTH_TEST)
	// gl.DepthMask(false)
	gl.DepthFunc(gl.LESS)
	gl.Enable(gl.CULL_FACE)
	gl.CullFace(gl.FRONT)
	// gl.FrontFace(gl.CW)
	glfw.SwapInterval(0) // vsync 0-off, 1-on
	var counter int = 0
	// for _, renderer := range renderers {
	// 	renderer.CalculateMesh()
	// }
	thisWorld.CalculateMesh()
	var currentTime float64

	for !window.ShouldClose() {

		currentTime = glfw.GetTime()
		deltaTime = currentTime - lastFrame
		lastFrame = currentTime

		processInput(window)

		gl.ClearColor(0.0, 0.6, 0.6, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		shader.Use()

		// shader.SetUniformMat4("projection\x00", mgl32.Perspective(mgl32.DegToRad(fov), float32(windowWidth)/float32(windowHeight), drawMinDistance, drawMaxDistance))
		shader.SetUniformMat4("camera\x00", camera.Mat4())
		// for _, renderer := range renderers {
		// 	renderer.Render(shader)
		// }
		thisWorld.Render(shader)

		// basicShader.Use()
		// basicShader.SetUniformMat4("camera\x00", camera.Mat4())
		// for _, l := range thisWorld.Lines() {
		// 	line.Render(basicShader, l[0], l[1])
		// }

		counter++
		if counter == 30 {
			window.SetTitle(fmt.Sprint("gocraft | FPS:", 1.0/deltaTime))
			counter = 0
			// PrintMemUsage()
		}

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func processInput(window *glfw.Window) {
	// glfw.GetKeyName(glfw.KeyEscape, glfw.GetKeyScancode(glfw.KeyEscape))
	// window.SetCharModsCallback()
	// glfw.GetCurrentContext().SetKeyCallback()
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		// fmt.Println("KeyEscape Down")
		window.SetShouldClose(true)
	}
	if !keyPressed && window.GetKey(glfw.Key1) == glfw.Press {
		// fmt.Println("Key1 Down")
		keyPressed = true
		switch drawMode {
		case gl.FILL:
			drawMode = gl.LINE
		case gl.LINE:
			drawMode = gl.FILL
		}
		gl.PolygonMode(gl.FRONT_AND_BACK, drawMode)
	}
	if keyPressed && window.GetKey(glfw.Key1) == glfw.Release {
		// fmt.Println("Key1 Up")
		keyPressed = false
	}

	if !keyPressed2 && window.GetKey(glfw.Key2) == glfw.Press {
		// fmt.Println("Key2 Down")
		keyPressed2 = true
		cursorHasDisabled = !cursorHasDisabled
		if cursorHasDisabled {
			window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
			camera.Unlock()
		} else {
			window.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
			camera.Lock()
		}
	}
	if keyPressed2 && window.GetKey(glfw.Key2) == glfw.Release {
		// fmt.Println("Key2 Up")
		keyPressed2 = false
	}

	if !keyPressed3 && window.GetKey(glfw.Key3) == glfw.Press {
		// fmt.Println("Key3 Down")
		keyPressed3 = true
		switch movementModel {
		case 0:
			movementModel = 1
			camera.SetMovementModel(control.ProvideCameraMovementModelFps())
		default:
			movementModel = 0
			camera.SetMovementModel(control.ProvideCameraMovementModelFree())
		}
	}
	if keyPressed3 && window.GetKey(glfw.Key3) == glfw.Release {
		// fmt.Println("Key3 Up")
		keyPressed3 = false
	}

	if window.GetKey(glfw.KeyW) == glfw.Press {
		// fmt.Println("KeyW Down")
		camera.MoveForward(speed * float32(deltaTime))
		// cameraPosition = cameraPosition.Add(cameraFront.Mul(speed * float32(deltaTime)))
	}

	if window.GetKey(glfw.KeyS) == glfw.Press {
		// fmt.Println("KeyS Down")
		camera.MoveBackward(speed * float32(deltaTime))
		// cameraPosition = cameraPosition.Sub(cameraFront.Mul(speed * float32(deltaTime)))
	}

	if window.GetKey(glfw.KeyA) == glfw.Press {
		// fmt.Println("KeyA Down")
		camera.MoveLeft(speed * float32(deltaTime))
		// cameraPosition = cameraPosition.Sub(cameraFront.Cross(cameraUp).Normalize().Mul(speed * float32(deltaTime)))
	}

	if window.GetKey(glfw.KeyD) == glfw.Press {
		// fmt.Println("KeyD Down")
		camera.MoveRight(speed * float32(deltaTime))
		// cameraPosition = cameraPosition.Add(cameraFront.Cross(cameraUp).Normalize().Mul(speed * float32(deltaTime)))
	}

	if window.GetKey(glfw.KeyLeftShift) == glfw.Press {
		// fmt.Println("KeyLeftShift Down")
		camera.MoveDown(speed * float32(deltaTime))
		// cameraPosition = cameraPosition.Sub(cameraUp.Mul(speed * float32(deltaTime)))
	}

	if window.GetKey(glfw.KeySpace) == glfw.Press {
		// fmt.Println("KeySpace Down")
		camera.MoveUp(speed * float32(deltaTime))
		// cameraPosition = cameraPosition.Add(cameraUp.Mul(speed * float32(deltaTime)))
	}
}

var lpmPressed, ppmPressed bool

func mouseButtonCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if !lpmPressed && button == glfw.MouseButtonLeft && action == glfw.Press {
		lpmPressed = true
		if block := thisWorld.FindBlockByRaycast(camera.Position(), camera.Direction(), 5); block != nil {
			block.DeleteFromWorld()
		}
	}

	if lpmPressed && button == glfw.MouseButtonLeft && action == glfw.Release {
		lpmPressed = false
	}
}

func cursorCallback(window *glfw.Window, xpos, ypos float64) {
	camera.MouseMoved(xpos, ypos)
}

func framebufferSizeCallback(window *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
	shader.Use()
	shader.SetUniformMat4("projection\x00", mgl32.Perspective(mgl32.DegToRad(fov), float32(width)/float32(height), drawMinDistance, drawMaxDistance))
	// basicShader.Use()
	// basicShader.SetUniformMat4("projection\x00", mgl32.Perspective(mgl32.DegToRad(fov), float32(width)/float32(height), drawMinDistance, drawMaxDistance))
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
