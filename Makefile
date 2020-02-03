.PHONY: dep

dep-bin:
	@go get -v github.com/veandco/go-sdl2/{sdl,img,mix,ttf}
dep:
	@dep install

