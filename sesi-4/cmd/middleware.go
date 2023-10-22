package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func TraceId() fiber.Handler {
	return func(c *fiber.Ctx) error {
		incoming := LogTracking{
			Message: "Incoming Request",
			Method:  string(c.Context().Method()),
			Uri:     string(c.Context().Path()),
			TraceId: c.Locals("X-TRACE-ID"),
		}

		log.Printf("%v method=%s uri=%s trace_id=%v",
			incoming.Message, incoming.Method, incoming.Uri, incoming.TraceId,
		)

		next := c.Next()

		outgoing := LogTracking{
			Message: "Finish Request",
			Method:  string(c.Context().Method()),
			Uri:     string(c.Context().Path()),
			TraceId: c.Locals("X-TRACE-ID"),
		}

		log.Printf("%v method=%s uri=%s trace_id=%v",
			outgoing.Message, outgoing.Method, outgoing.Uri, outgoing.TraceId,
		)

		return next
	}
}
