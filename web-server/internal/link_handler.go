package internal

import (
	"github.com/gofiber/fiber/v2"
	"net/url"
)

type (
	CreateLinkRequest struct {
		External string `json:"external"`
		Internal string `json:"internal"`
	}

	GetLinkResponse struct {
		Internal string `json:"internal"`
	}
)

type NotFoundError struct{}

func (e NotFoundError) Error() string {
	return "Value is not found"
}

var links = make(map[string]string)

type LinkStorage interface {
	AddLink(request *CreateLinkRequest) error
	GetLink(external string) (error, string)
}

type LinkStorageMap struct{}

func (storage *LinkStorageMap) AddLink(request *CreateLinkRequest) error {
	links[request.External] = request.Internal
	return nil
}

func (storage *LinkStorageMap) GetLink(external string) (error, string) {
	if val, found := links[external]; found {
		return nil, val
	} else {
		return NotFoundError{}, ""
	}
}

type LinkHandler struct {
	Storage LinkStorage
}

func (h *LinkHandler) AddLink(c *fiber.Ctx) error {
	request := CreateLinkRequest{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
	}

	h.Storage.AddLink(&request)
	return c.SendStatus(fiber.StatusOK)
}

func (h *LinkHandler) GetLink(request string, c *fiber.Ctx) error {
	external, err := url.QueryUnescape(request)
	if err != nil {
		return err
	}

	if err, val := h.Storage.GetLink(external); err == nil {
		return c.Status(fiber.StatusOK).JSON(GetLinkResponse{val})
	} else {
		return c.Status(fiber.StatusNotFound).SendString("Link not found")
	}
}
