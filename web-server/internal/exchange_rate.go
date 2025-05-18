package internal

var exchangeRate = map[string]float64{
	"USD/EUR": 0.8,
	"EUR/USD": 1.25,
	"USD/GBP": 0.7,
	"GBP/USD": 1.43,
	"USD/JPY": 110,
	"JPY/USD": 0.0091,
}

type (
	BinarySearchRequest struct {
		Numbers []int `json:"numbers"`
		Target  int   `json:"target"`
	}

	BinarySearchResponse struct {
		TargetIndex int    `json:"target_index"`
		Error       string `json:"error,omitempty"`
	}
)

const targetNotFound = -1

//func main() {
//	webApp := fiber.New(fiber.Config{
//		ReadBufferSize: 16 * 1024})
//	webApp.Get("/", func(c *fiber.Ctx) error {
//		return c.SendStatus(200)
//	})
//
//	webApp.Post("/search", func(c *fiber.Ctx) error {
//		request := BinarySearchRequest{}
//		if err := c.BodyParser(&request); err != nil {
//			return c.Status(fiber.StatusBadRequest).JSON(BinarySearchResponse{targetNotFound, "Invalid JSON"})
//		}
//
//		targetIndex := slices.Index(request.Numbers, request.Target)
//		if targetIndex == targetNotFound {
//			return c.Status(fiber.StatusNotFound).JSON(BinarySearchResponse{targetNotFound, "Target was not found"})
//		} else {
//			return c.Status(fiber.StatusOK).JSON(BinarySearchResponse{targetIndex, ""})
//		}
//	})
//
//	logrus.Fatal(webApp.Listen(":8080"))
//}
