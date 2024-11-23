import axios from 'axios'
import { toCapitalCase } from './utils'

export const BACKEND_URL = import.meta.env.VITE_BACKEND_URL ?? 'localhost:3000'

export type Condition = {
  additional_info?: string
  age?: number
  code?: string
  country?: string
  date_range?: {
    from?: string
    to?: string
  }
  gender?: string
  locality?: string
  region?: string
  sport?: string
  event_type?: string
  event_scale?: string
}

export type Competition = {
  code: string
  start_date: string
  location_data: [
    {
      country: string
      region: string
      locality: string
    },
  ]
  age_data: [
    {
      gender: string
      left_bound: number
      right_bound: number
      original: string
    },
  ]
  title: string
  additional_info: string
  n_participants: number
  stage: string
  end_date: string
  sport: string
}

type FilterResult = {
  events: Competition[]
  total: number
}

// axios.interceptors.request.use(request => {
//   console.log('Starting Request', JSON.stringify(request, null, 2))
//   return request
// })

export type SubscriptionRequest = {
  additional_info?: string
  age?: number
  code?: string
  country?: string
  email: string
  end_date?: string
  event_scale?: string
  event_type?: string
  gender?: string
  locality?: string
  region?: string
  sport?: string
  stage?: string
  start_date?: string
}

export let getEvents = async (
  page: number,
  page_size: number,
  condition: Condition = {},
  required_fields: string[] = [],
): Promise<FilterResult> =>
  await axios
    .post(BACKEND_URL + '/filter', {
      condition,
      pagination: {
        page,
        page_size,
      },
      required_fields,
    })
    .then((res) => res.data)

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
