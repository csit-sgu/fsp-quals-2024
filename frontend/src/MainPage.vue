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

import { sports, countries, getRegions, getLocalities, countryHasRegions } from '@/lib/dataSource'
import SidebarFooter from './components/ui/sidebar/SidebarFooter.vue'
import { type Competition, getEvents } from '@/lib/dataSource'
import { Toaster } from '@/components/ui/sonner'

const route = useRoute()
const router = useRouter()

const getPageSize = () => (route.path == '/weekly' ? 10 : 25)

const updateViewMode = (newValue: string) => {
  router.push(newValue)
}

const pickedSport = ref('')
const updateSport = (newValue: string) => (pickedSport.value = newValue)

const pickedCountry = ref('')
const pickedCountryRegions = ref([])

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

const pickedRegion = ref('')
const pickedRegionLocalities = ref([])

const updateRegion = async (newValue: string) => {
  pickedRegion.value = newValue
  pickedRegionLocalities.value = await getLocalities(pickedCountry.value, newValue)
  pickedLocality.value = ''
}

const pickedLocality = ref('')
const updateLocality = (newValue: string) => (pickedLocality.value = newValue)

const discipline = ref('')
const disciplineUpdated = (newValue: string | number) => (discipline.value = newValue.toString())

const additionalInfo = ref('')
const additionalInfoUpdated = (newValue: string | number) =>
  (additionalInfo.value = newValue.toString())

type EventsCounter = {
  events: Competition[]
  count: number
}

const eventsWithCount: Ref<EventsCounter> = ref({ events: [], count: 0 })
const total: Ref<number> = ref(1)
onMounted(async () => {
  const res = await getEvents(0, getPageSize())
  eventsWithCount.value.events = res.events
  eventsWithCount.value.count++
  total.value = res.total
})
const updateEvents = async (page: number) => {
  const res = await getEvents(page, getPageSize())
  eventsWithCount.value.events = res.events
  eventsWithCount.value.count++
  total.value = res.total
  window.scrollTo(0, 0)
}

const getFilters = () => {
  return [
    pickedSport.value,
    pickedCountry.value,
    pickedRegion.value,
    pickedLocality.value,
    discipline.value + ' ' + additionalInfo.value,
  ]
}

const applyFilters = () => {
  // TODO(mchernigin): implement filters
  console.log(getFilters())
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
            <DateFilterPicker @update="updateViewMode" />

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
              @update="updateLocality"
            />

            <SidebarGroupLabel class="pt-8 pb-6">Фильтрация по соревнованиям</SidebarGroupLabel>
            <Chooser
              :options="sports"
              :show-search="true"
              default-msg="Любой вид спорта"
              @update="updateSport"
            />
            <div class="pt-2">
              <Input
                @update:model-value="disciplineUpdated"
                type="search"
                placeholder="Любая дисциплина..."
                class="pt-2"
              />
            </div>
            <div class="pt-2">
              <Input
                @update:model-value="additionalInfoUpdated"
                type="search"
                placeholder="Поиск по другой информации..."
                class="pt-2"
              />
            </div>
          </SidebarGroup>
        </ScrollArea>
      </SidebarContent>
      <SidebarFooter class="p-8">
        <Button @click="applyFilters">Применить фильтры</Button>
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
        v-if="route.path === '/table'"
      />
      <Pagination
        :key="total"
        v-if="eventsWithCount.events.length > 0"
        v-slot="{ page }"
        :itemsPerPage="getPageSize()"
        :total="total"
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

        <div class="pt-4 color-red-500 text-center">Всего найдено соревнований: {{ total }}</div>
      </Pagination>
    </SidebarInset>
  </SidebarProvider>
</template>
