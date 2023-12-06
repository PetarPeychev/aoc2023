package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Almanac struct {
	Seeds                 []int
	SeedToSoil            []MapRange
	SoilToFertilizer      []MapRange
	FertilizerToWater     []MapRange
	WaterToLight          []MapRange
	LightToTemperature    []MapRange
	TemperatureToHumidity []MapRange
	HumidityToLocation    []MapRange
}

func convert(almanac Almanac, mapRanges []MapRange, input int) int {
	for _, mapRange := range mapRanges {
		if input >= mapRange.SourceStart && input < mapRange.SourceStart+mapRange.Length {
			return mapRange.DestinationStart + input - mapRange.SourceStart
		}
	}
	return input
}

type MapRange struct {
	SourceStart      int
	DestinationStart int
	Length           int
}

func main() {
	file, _ := os.Open("input1.txt")
	defer file.Close()

	line := 0
	almanac := Almanac{}

	const (
		SEED_TO_SOIL = iota
		SOIL_TO_FERTILIZER
		FERTILIZER_TO_WATER
		WATER_TO_LIGHT
		LIGHT_TO_TEMPERATURE
		TEMPERATURE_TO_HUMIDITY
		HUMIDITY_TO_LOCATION
	)

	currentBlock := -1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line++
		text := scanner.Text()
		if line == 1 {
			seeds := strings.Split(text[6:], " ")
			for _, seed := range seeds {
				if seed != "" {
					seedNum, _ := strconv.Atoi(seed)
					almanac.Seeds = append(almanac.Seeds, seedNum)
				}
			}
		} else if strings.Contains(text, "seed-to-soil") {
			currentBlock = SEED_TO_SOIL
		} else if strings.Contains(text, "soil-to-fertilizer") {
			currentBlock = SOIL_TO_FERTILIZER
		} else if strings.Contains(text, "fertilizer-to-water") {
			currentBlock = FERTILIZER_TO_WATER
		} else if strings.Contains(text, "water-to-light") {
			currentBlock = WATER_TO_LIGHT
		} else if strings.Contains(text, "light-to-temperature") {
			currentBlock = LIGHT_TO_TEMPERATURE
		} else if strings.Contains(text, "temperature-to-humidity") {
			currentBlock = TEMPERATURE_TO_HUMIDITY
		} else if strings.Contains(text, "humidity-to-location") {
			currentBlock = HUMIDITY_TO_LOCATION
		} else if text == "" {
			continue
		} else {
			numStrs := strings.Split(text, " ")
			var nums []int
			for _, num := range numStrs {
				if num != "" {
					num, _ := strconv.Atoi(num)
					nums = append(nums, num)
				}
			}
			switch currentBlock {
			case SEED_TO_SOIL:
				almanac.SeedToSoil = append(almanac.SeedToSoil, MapRange{nums[1], nums[0], nums[2]})
			case SOIL_TO_FERTILIZER:
				almanac.SoilToFertilizer = append(almanac.SoilToFertilizer, MapRange{nums[1], nums[0], nums[2]})
			case FERTILIZER_TO_WATER:
				almanac.FertilizerToWater = append(almanac.FertilizerToWater, MapRange{nums[1], nums[0], nums[2]})
			case WATER_TO_LIGHT:
				almanac.WaterToLight = append(almanac.WaterToLight, MapRange{nums[1], nums[0], nums[2]})
			case LIGHT_TO_TEMPERATURE:
				almanac.LightToTemperature = append(almanac.LightToTemperature, MapRange{nums[1], nums[0], nums[2]})
			case TEMPERATURE_TO_HUMIDITY:
				almanac.TemperatureToHumidity = append(almanac.TemperatureToHumidity, MapRange{nums[1], nums[0], nums[2]})
			case HUMIDITY_TO_LOCATION:
				almanac.HumidityToLocation = append(almanac.HumidityToLocation, MapRange{nums[1], nums[0], nums[2]})
			}
		}
	}
	fmt.Printf("%+v\n", almanac)

	minLocation := 1000000000
	for _, seed := range almanac.Seeds {
		soil := convert(almanac, almanac.SeedToSoil, seed)
		fertilizer := convert(almanac, almanac.SoilToFertilizer, soil)
		water := convert(almanac, almanac.FertilizerToWater, fertilizer)
		light := convert(almanac, almanac.WaterToLight, water)
		temperature := convert(almanac, almanac.LightToTemperature, light)
		humidity := convert(almanac, almanac.TemperatureToHumidity, temperature)
		location := convert(almanac, almanac.HumidityToLocation, humidity)
		fmt.Printf("seed: %d, soil: %d, fertilizer: %d, water: %d, light: %d, temperature: %d, humidity: %d, location: %d\n", seed, soil, fertilizer, water, light, temperature, humidity, location)
		if location < minLocation {
			minLocation = location
		}
	}
	fmt.Printf("Closest seed: %d\n", minLocation)
}
