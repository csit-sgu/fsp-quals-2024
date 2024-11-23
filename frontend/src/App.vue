<script setup lang="ts">
import { useRoute } from 'vue-router'
import { Separator } from '@/components/ui/separator'
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupLabel,
  SidebarInset,
  SidebarProvider,
  SidebarTrigger,
} from '@/components/ui/sidebar'
import { WeeklyView, DateFilterPicker, Chooser } from '@/components/ui'
import { Button } from '@/components/ui/button'
import { ref } from 'vue'

import { sports, countries, getRegions, getLocalities, countryHasRegions } from '@/lib/dataSource'
import SidebarFooter from './components/ui/sidebar/SidebarFooter.vue'

const route = useRoute()

const showMailSubscriptionDialog = ref(false)

const viewMode = ref('')
const updateViewMode = (newValue: string) => (viewMode.value = newValue)

const pickedSport = ref('')
const updateSport = (newValue: string) => (pickedSport.value = newValue)

const pickedCountry = ref('')
const pickedCountryRegions = ref([])

const updateCountry = async (newValue: string) => {
  pickedCountry.value = newValue
  if (newValue == 'РОССИЯ') {
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
</script>

<template>
  <SidebarProvider>
    <Sidebar collapsible="offcanvas" class="sticky top-0 h-screen">
      <SidebarContent>
        <SidebarGroup class="content-center px-4 w-auto">
          <SidebarGroupLabel class="pt-4 pb-6">Фильтрация по дате</SidebarGroupLabel>
          <DateFilterPicker @update="updateViewMode" />
          <SidebarGroupLabel class="pt-8 pb-6">Фитрация по соревнованиям</SidebarGroupLabel>
          <Chooser :options="sports" :show-search="true" default-msg="Любой вид спорта" @update="updateSport" />
          <SidebarGroupLabel class="pt-8 pb-6">Фитрация по месту проведения</SidebarGroupLabel>
          <Chooser :options="countries" :show-search="true" default-msg="Любая страна" @update="updateCountry" />
          <div v-if="countryHasRegions(pickedCountry)" class="pb-2" />
          <Chooser v-if="countryHasRegions(pickedCountry)" :show-search="true" :options="pickedCountryRegions"
            default-msg="Любой регион" @update="updateRegion" />
          <div class="pt-2" />
          <Chooser v-if="
            pickedRegion.length > 0 ||
            (!countryHasRegions(pickedCountry) && pickedCountry.length > 0)
          " :show-search="true" :options="pickedRegionLocalities" default-msg="Любой населённый пункт"
            @update="updateLocality" />
        </SidebarGroup>
      </SidebarContent>
      <SidebarFooter>
        <Button @click="updateViewMode('table')">Применить фильтры</Button>
        <Button variant="outline" @click="showMailSubscriptionDialog = true">Подписаться на уведомления по
          фильтрам</Button>
      </SidebarFooter>
    </Sidebar>
    <SidebarInset class="min-h-screen overflow-x-hidden">
      <header
        class="flex w-full h-16 shrink-0 items-center gap-2 transition-[width,height] ease-linear group-has-[[data-collapsible=icon]]/sidebar-wrapper:h-12">
        <div class="flex items-center gap-2 px-4">
          <SidebarTrigger class="-ml-1" />
          <Separator orientation="vertical" class="mr-2 h-4" />
        </div>
      </header>
      <WeeklyView v-if="route.path === '/weekly'" />
      <WeeklyView v-if="route.path === '/table'" />
    </SidebarInset>
  </SidebarProvider>
</template>
