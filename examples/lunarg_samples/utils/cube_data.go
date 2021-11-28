package utils

type Vertex struct {
	PosX, PosY, PosZ, PosW float32
	R, G, B, A             float32
}

type VertexUV struct {
	PosX, PosY, PosZ, PosW float32
	U, V                   float32
}

//var VBData = []Vertex{
//	{-1, -1, -1, 1, 0, 0, 0, 1},
//	{1, -1, -1, 1, 1, 0, 0, 1},
//	{-1, 1, -1, 1, 0, 1, 0, 1},
//	{-1, 1, -1, 1, 0, 1, 0, 1},
//	{1, -1, -1, 1, 1, 0, 0, 1},
//	{1, 1, -1, 1, 1, 1, 0, 1},
//	{-1, -1, 1, 1, 0, 0, 1, 1},
//}

var VBTextureData = []VertexUV{
	// left face
	{-1, -1, -1, 1, 1, 0},
	{-1, 1, 1, 1, 0, 1},
	{-1, -1, 1, 1, 0, 0},
	{-1, 1, 1, 1, 0, 1},
	{-1, -1, -1, 1, 1, 0},
	{-1, 1, -1, 1, 1, 1},

	// front face
	{-1, -1, -1, 1, 0, 0},
	{1, -1, -1, 1, 1, 0},
	{1, 1, -1, 1, 1, 1},
	{-1, -1, -1, 1, 0, 0},
	{1, 1, -1, 1, 1, 1},
	{-1, 1, -1, 1, 0, 1},

	// top face
	{-1, -1, -1, 1, 0, 1},
	{1, -1, 1, 1, 1, 0},
	{1, -1, -1, 1, 1, 1},
	{-1, -1, -1, 1, 0, 1},
	{-1, -1, 1, 1, 0, 0},
	{1, -1, 1, 1, 1, 0},

	// bottom face
	{-1, 1, -1, 1, 0, 0},
	{1, 1, 1, 1, 1, 1},
	{-1, 1, 1, 1, 0, 1},
	{-1, 1, -1, 1, 0, 0},
	{1, 1, -1, 1, 1, 0},
	{1, 1, 1, 1, 1, 1},

	// right face
	{1, 1, -1, 1, 0, 1},
	{1, -1, 1, 1, 1, 0},
	{1, 1, 1, 1, 1, 1},
	{1, -1, 1, 1, 1, 0},
	{1, 1, -1, 1, 0, 1},
	{1, -1, -1, 1, 0, 0},

	// back face
	{-1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 0, 1},
	{-1, -1, 1, 1, 1, 0},
	{-1, -1, 1, 1, 1, 0},
	{1, 1, 1, 1, 0, 1},
	{1, -1, 1, 1, 0, 0},
}
