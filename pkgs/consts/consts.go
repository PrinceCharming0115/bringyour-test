package consts

import (
	msg "bringyour-test/pkgs/models"
	"math/rand"
	"strings"
)

const (
	MockUUID = "########-####-####-####-############"
)

func ShortMessage(message msg.Message) string {
	if message.Prefix == "message" && message.UUID == MockUUID {
		return "message X"
	} else if message.Prefix == "message" && message.UUID != MockUUID {
		return "message Y"
	} else if message.Prefix == "ok" && message.UUID == MockUUID {
		return "ok X"
	} else if message.Prefix == "ok" && message.UUID != MockUUID {
		return "ok Y"
	}
	return strings.ToUpper(message.Prefix)
}

func RandomMessage() string {
	var messages = []string{
		"Simplicity is the ultimate sophistication.",
		"The noblest pleasure is the joy of understanding.",
		"Realize that everything connects to everything else.",
		"Learning never exhausts the mind.",
		"The greatest deception men suffer is from their own opinions.",
		"Intellectual passion drives out sensuality.",
		"Small steps are the fastest path to achieving great goals.",
		"It had long since come to my attention that people of accomplishment rarely sat back and let things happen to them. They went out and happened to things.",
		"Obstacles cannot crush me. Every obstacle yields to stern resolve.",
		"I have been impressed with the urgency of doing. Knowing is not enough; we must apply. Being willing is not enough; we must do.",
		"Art is never finished, only abandoned.",
		"The greatest geniuses sometimes accomplish more when they work less.",
		"The function of muscle is to pull, not to push, except in the case of the genitals and the tongue.",
		"The human foot is a masterpiece of engineering and a work of art.",
		"I love those who can smile in trouble, who can gather strength from distress, and grow brave by reflection.",
		"Time stays long enough for those who use it.",
		"Painting is poetry that is seen rather than felt, and poetry is painting that is felt rather than seen.",
		"The human bird shall take his first flight, filling the world with amazement, all writings with his fame, and bringing eternal glory to the nest whence he sprang.",
		"Dwell on the beauty of life. Watch the stars, and see yourself running with them.",
		"You can have no dominion greater or less than that over yourself.",
		"The most beautiful thing we can experience is the mysterious. It is the source of all true art and all science.",
		"Imagination is more important than knowledge.",
		"The only source of knowledge is experience.",
		"The only way to do great work is to love what you do.",
		"Everything should be made as simple as possible, but not simpler.",
		"The important thing is not to stop questioning. Curiosity has its own reason for existing.",
		"The true sign of intelligence is not knowledge but imagination.",
		"In the middle of difficulty lies opportunity.",
		"The most beautiful experience we can have is the mysterious.",
		"I have no special talent. I am only passionately curious.",
		"The only limit to our realization of tomorrow will be our doubts of today.",
		"The world as we have created it is a process of our thinking. It cannot be changed without changing our thinking.",
		"Try not to become a man of success, but rather try to become a man of value.",
		"The measure of intelligence is the ability to change.",
		"Education is what remains after one has forgotten what one has learned in school.",
		"The person who follows the crowd will usually go no further than the crowd. The person who walks alone is likely to find himself in places no one has ever been.",
		"Two things are infinite: the universe and human stupidity; and I'm not sure about the universe.",
		"Insanity is doing the same thing over and over again and expecting different results.",
		"Look deep into nature, and then you will understand everything better.",
		"The most incomprehensible thing about the universe is that it is comprehensible.",
	}
	return messages[rand.Intn(len(messages))]
}
