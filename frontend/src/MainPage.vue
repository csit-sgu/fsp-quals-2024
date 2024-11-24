<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router'
import { Separator } from '@/components/ui/separator'
import {
  Pagination,
  PaginationList,
  PaginationListItem,
  PaginationFirst,
  PaginationLast,
  PaginationNext,
  PaginationPrev,
  PaginationEllipsis,
} from '@/components/ui/pagination'
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupLabel,
  SidebarInset,
  SidebarProvider,
  SidebarTrigger,
} from '@/components/ui/sidebar'
import { WeeklyView, DateFilterPicker, Chooser, TableView, SubscribeDialog } from '@/components/ui'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { onMounted, ref, type Ref } from 'vue'
import { ScrollArea } from '@/components/ui/scroll-area'
import dayjs from 'dayjs'
import isoWeek from 'dayjs/plugin/isoWeek'
import utc from 'dayjs/plugin/utc'
import { scaleMap, sexMap, typeMap } from '@/lib/dataSource'
import {
  sports,
  countries,
  getRegions,
  getLocalities,
  countryHasRegions,
  type Condition,
} from '@/lib/dataSource'
import SidebarFooter from './components/ui/sidebar/SidebarFooter.vue'
import { type Competition, getEvents } from '@/lib/dataSource'
import { Toaster } from '@/components/ui/sonner'

const route = useRoute()
const router = useRouter()

const getPageSize = () => (route.path == '/weekly' ? 10 : 25)

const viewMode = ref('')
const calendarRange: Ref<any> = ref(null)

const pickedSport = ref('')

const pickedCountry = ref('')
const pickedCountryRegions = ref([])
const pickedRegion = ref('')
const pickedRegionLocalities = ref([])
const pickedLocality = ref('')

const title = ref('')
const discipline = ref('')
const additionalInfo = ref('')

const gender = ref('')
const age = ref(0)

const event_type = ref('')
const event_scale = ref('')

const updateViewMode = (path: string, value: string) => {
  router.push(path)
  viewMode.value = value
  updateEvents(0)
}

const updateCountry = async (newValue: string) => {
  pickedCountry.value = newValue
  if (countryHasRegions(newValue)) {
    pickedCountryRegions.value = await getRegions(newValue)
  } else {
    pickedCountryRegions.value = []
    pickedRegionLocalities.value = await getLocalities(newValue, '')
  }
  pickedRegion.value = ''
  pickedLocality.value = ''
}

const updateRegion = async (newValue: string) => {
  pickedRegion.value = newValue
  pickedRegionLocalities.value = await getLocalities(pickedCountry.value, newValue)
  pickedLocality.value = ''
}

const eventsWithCount: Ref<{
  events: Competition[]
  count: number
}> = ref({ events: [], count: 0 })
const total_items: Ref<number> = ref(1)
const updateEvents = async (page: number) => {
  const res = await getEvents(page, getPageSize(), getFilters())
  eventsWithCount.value.events = res.events
  eventsWithCount.value.count++
  total_items.value = res.total
  window.scrollTo(0, 0)
}
onMounted(async () => updateEvents(0))

dayjs.extend(isoWeek)
dayjs.extend(utc)

const getFilters = (): Condition => {
  const date_range = () => {
    const from = dayjs.utc().format('YYYY-MM-DD')
    if (viewMode.value === 'week') {
      return { from, to: dayjs.utc().add(6, 'day').format('YYYY-MM-DD') }
    }
    if (viewMode.value === 'month') {
      return { from, to: dayjs.utc().add(1, 'month').format('YYYY-MM-DD') }
    }
    if (viewMode.value === 'quarter') {
      return { from, to: dayjs.utc().add(3, 'month').format('YYYY-MM-DD') }
    }
    if (viewMode.value === 'half') {
      return { from, to: dayjs.utc().add(6, 'month').format('YYYY-MM-DD') }
    }
    if (
      viewMode.value === 'custom' &&
      calendarRange.value &&
      calendarRange.value.start &&
      calendarRange.value.end
    ) {
      return { from: calendarRange.value.start.toString(), to: calendarRange.value.end.toString() }
    }
  }

  const additional_info = () => {
    let s = ''
    s += discipline.value
    s += '|'
    s += additionalInfo.value
    return s.trim()
  }

  return {
    title: title.value,
    additional_info: additional_info(),
    sport: pickedSport.value.toUpperCase(),
    age: age.value,
    // code?: string, // NOTE(mchernigin): not used for filters
    country: pickedCountry.value,
    region: pickedRegion.value,
    locality: pickedLocality.value,
    date_range: date_range(),
    gender: sexMap.get(gender.value),
    event_type: typeMap.get(event_type.value),
    event_scale: scaleMap.get(event_scale.value),
  }
}
</script>

<template>
  <Toaster richColors />
  <SidebarProvider>
    <Sidebar collapsible="offcanvas" class="sticky top-0 h-screen">
      <SidebarContent>
        <ScrollArea>
          <SidebarGroup class="content-center px-4 w-auto">
            <h1 class="text-2xl font-extrabold px-2 py-4">Поиск соревнований</h1>
            <SidebarGroupLabel class="pt-4 pb-6">Фильтрация по дате</SidebarGroupLabel>
            <DateFilterPicker
              @update="updateViewMode"
              @calendarUpdate="
                (range: any) => {
                  calendarRange = range
                }
              "
            />

            <SidebarGroupLabel class="pt-8 pb-6">Фильтрация по месту проведения</SidebarGroupLabel>
            <Chooser
              :options="countries"
              :show-search="true"
              default-msg="Любая страна"
              @update="updateCountry"
            />
            <div v-if="countryHasRegions(pickedCountry)" class="pb-2" />
            <Chooser
              v-if="countryHasRegions(pickedCountry)"
              :show-search="true"
              :options="pickedCountryRegions"
              default-msg="Любой регион"
              @update="updateRegion"
            />
            <div class="pt-2" />
            <Chooser
              v-if="
                pickedRegion.length > 0 ||
                (!countryHasRegions(pickedCountry) && pickedCountry.length > 0)
              "
              :show-search="true"
              :options="pickedRegionLocalities"
              default-msg="Любой населённый пункт"
              @update="(newValue: string) => (pickedLocality = newValue)"
            />

            <SidebarGroupLabel class="pt-6 pb-6">Фильтрация по соревнованиям</SidebarGroupLabel>
            <Chooser
              :options="sports"
              :show-search="true"
              default-msg="Любой вид спорта"
              @update="(newValue: string) => (pickedSport = newValue)"
            />
            <div class="pt-2">
              <Input
                @update:model-value="
                  (newValue: string | number) => (discipline = newValue.toString())
                "
                type="search"
                placeholder="Поиск по дисциплине..."
              />
            </div>
            <div class="pt-2">
              <Chooser
                :options="[...typeMap.keys()]"
                :show-search="true"
                default-msg="Любой тип соревнования"
                @update="(newValue: string) => (event_type = String(newValue))"
              />
            </div>
            <div class="pt-2">
              <Chooser
                :options="[...scaleMap.keys()]"
                default-msg="Любой уровень соревнования"
                @update="(newValue: string) => (event_scale = String(newValue))"
              />
            </div>
            <div class="pt-2">
              <Input
                @update:model-value="(newValue: string | number) => (title = newValue.toString())"
                type="search"
                placeholder="Поиск по названию..."
              />
            </div>
            <div class="pt-2">
              <Input
                @update:model-value="
                  (newValue: string | number) => (additionalInfo = String(newValue))
                "
                type="search"
                placeholder="Поиск по другой информации..."
              />
            </div>

            <SidebarGroupLabel class="pt-10 pb-8"
              >Фильтрация по информации о спортсмене</SidebarGroupLabel
            >
            <Chooser
              :options="['Мужской', 'Женский']"
              default-msg="Любой пол"
              @update="(newValue: string) => (gender = String(newValue))"
            />
            <div class="pt-2">
              <Input
                @update:model-value="(newValue: string | number) => (age = Number(newValue))"
                type="number"
                placeholder="Любой возраст..."
              />
            </div>
          </SidebarGroup>
        </ScrollArea>
      </SidebarContent>
      <SidebarFooter class="p-8">
        <Button @click="updateEvents(0)">Применить фильтры</Button>
        <SubscribeDialog :searchFilters="getFilters()" />
      </SidebarFooter>
    </Sidebar>
    <SidebarInset class="min-h-screen overflow-x-hidden">
      <header
        class="flex w-full h-16 shrink-0 items-center gap-2 transition-[width,height] ease-linear group-has-[[data-collapsible=icon]]/sidebar-wrapper:h-12"
      >
        <div class="flex items-center gap-2 px-4">
          <SidebarTrigger class="-ml-1" />
          <Separator orientation="vertical" class="mr-2 h-4" />
        </div>
      </header>
      <WeeklyView
        :key="eventsWithCount.count"
        :events="eventsWithCount.events"
        v-if="route.path === '/weekly'"
      />
      <TableView
        :key="eventsWithCount.count"
        :events="eventsWithCount.events"
        v-if="route.path === '/' || route.path === '/table'"
      />
      <Pagination
        :key="total_items"
        v-if="eventsWithCount.events && eventsWithCount.events.length > 0"
        v-slot="{ page }"
        :itemsPerPage="getPageSize()"
        :total="total_items - getPageSize()"
        :sibling-count="1"
        show-edges
        :default-page="1"
        class="self-center p-16"
      >
        <PaginationList v-slot="{ items }" class="flex items-center gap-1">
          <PaginationFirst />
          <PaginationPrev />

          <template v-for="(item, index) in items">
            <PaginationListItem
              v-if="item.type === 'page'"
              :key="index"
              :value="item.value"
              as-child
            >
              <Button
                class="w-10 h-10 p-0"
                :variant="item.value === page ? 'default' : 'outline'"
                @click="updateEvents(item.value)"
              >
                {{ item.value }}
              </Button>
            </PaginationListItem>
            <PaginationEllipsis v-else :key="item.type" :index="index" />
          </template>

          <PaginationNext />
          <PaginationLast />
        </PaginationList>

        <div class="pt-4 color-red-500 text-center">
          Всего найдено соревнований: {{ total_items }}
        </div>
      </Pagination>
    </SidebarInset>
  </SidebarProvider>
</template>
