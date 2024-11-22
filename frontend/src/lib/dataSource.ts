import axios from 'axios'

export const sports = [
  'Дартс зубочистками',
  'Бег по углям',
  'Бег от радости',
  'Зимнее программирование',
]

const BACKEND_URL = import.meta.env.VITE_BACKEND_URL ?? "localhost:3000"

export const countries = await axios.get(BACKEND_URL + '/countries').then(res => res.data)

export async function getRegions(country: string) {
  return await axios.get(BACKEND_URL + '/regions', { params: { country } }).then(res => res.data)
}

export async function getLocalities(country: string, region: string) {
  return await axios.get(BACKEND_URL + '/localities', { params: { country, region } }).then(res => res.data)
}

