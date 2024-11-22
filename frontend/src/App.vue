<script setup lang="ts">
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
import { ref } from 'vue'

import { sports, countries, getRegions, getLocalities } from '@/lib/dataSource'

const viewMode = ref('')
const updateViewMode = (newValue: string) => (viewMode.value = newValue)

const pickedSport = ref('')
const updateSport = (newValue: string) => (pickedSport.value = newValue)

const pickedCountry = ref('')
const pickedCountryRegions = ref([])

const updateCountry = async (newValue: string) => {
  pickedCountry.value = newValue
  pickedCountryRegions.value = await getRegions(newValue)
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
    <Sidebar collapsible="offcanvas">
      <SidebarContent>
        <SidebarGroup class="content-center px-4 w-auto">
          <SidebarGroupLabel class="pt-4 pb-6">Фильтрация по дате</SidebarGroupLabel>
          <DateFilterPicker @update="updateViewMode" />
          <SidebarGroupLabel class="pt-8 pb-6">Фитрация по соревнованиям</SidebarGroupLabel>
          <Chooser :options="sports" :show-search="true" default-msg="Любой вид спорта" @update="updateSport" />
          <SidebarGroupLabel class="pt-8 pb-6">Фитрация по месту проведения</SidebarGroupLabel>
          <Chooser :options="countries" :show-search="true" default-msg="Любая страна" @update="updateCountry" />
          <div v-if="pickedCountry == 'Россия'" class="pb-2" />
          <Chooser v-if="pickedCountry == 'Россия'" :show-search="true" :options="pickedCountryRegions"
            default-msg="Любой регион" @update="updateRegion" />
          <div class="pt-2" />
          <Chooser v-if="pickedRegion.length > 0 || (pickedCountry != 'Россия' && pickedCountry.length > 0)"
            :show-search="true" :options="pickedRegionLocalities" default-msg="Любой населённый пункт"
            @update="updateLocality" />
        </SidebarGroup>
      </SidebarContent>
    </Sidebar>
    <SidebarInset class="min-h-screen overflow-x-hidden">
      <header
        class="flex w-full h-16 shrink-0 items-center gap-2 transition-[width,height] ease-linear group-has-[[data-collapsible=icon]]/sidebar-wrapper:h-12">
        <div class="flex items-center gap-2 px-4">
          <SidebarTrigger class="-ml-1" />
          <Separator orientation="vertical" class="mr-2 h-4" />
        </div>
      </header>
      <WeeklyView v-if="!viewMode || viewMode == 'Ближайшая неделя'" />
    </SidebarInset>
  </SidebarProvider>
</template>
