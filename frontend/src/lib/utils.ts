import { type ClassValue, clsx } from 'clsx'
import { twMerge } from 'tailwind-merge'

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export const toCapitalCase = (x: string) => x.charAt(0).toUpperCase() + x.slice(1).toLowerCase();
