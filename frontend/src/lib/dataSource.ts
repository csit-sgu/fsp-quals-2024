import axios from 'axios'
import { toCapitalCase } from './utils'

const BACKEND_URL = import.meta.env.VITE_BACKEND_URL ?? 'localhost:3000'

export const countryHasRegions = (country: string) => country == 'РОССИЯ'

export let sports = await axios
  .get(BACKEND_URL + '/sports')
  .then((res) => res.data)
  .catch((r) => {
    console.log(r)
    return []
  })
if (!Array.isArray(sports)) {
  sports = []
}
sports = sports.map(toCapitalCase)

export let countries = await axios
  .get(BACKEND_URL + '/countries')
  .then((res) => res.data)
  .catch((r) => {
    console.log(r)
    return []
  })
if (!Array.isArray(countries)) {
  countries = []
}

export async function getRegions(country: string) {
  return await axios
    .get(BACKEND_URL + '/regions', { params: { country } })
    .then((res) => res.data)
    .catch((r) => {
      console.log(r)
      return []
    })
}

export async function getLocalities(country: string, region: string) {
  return await axios
    .get(BACKEND_URL + '/localities', { params: { country, region } })
    .then((res) => res.data)
    .catch((r) => {
      console.log(r)
      return []
    })
}
