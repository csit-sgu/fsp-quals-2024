<script setup lang="ts">
import { Separator } from '@/components/ui/separator'
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupLabel,
  SidebarInset,
  SidebarProvider,
  SidebarRail,
  SidebarTrigger,
} from '@/components/ui/sidebar'
import { WeeklyView, DateFilterPicker, Chooser } from '@/components/ui'
import { ref } from 'vue'

import { sports } from '@/lib/dataSource'

const pickedSport = ref('')
const updateSport = (newValue: string) => pickedSport.value = newValue
</script>

<template>
  <SidebarProvider>
    <Sidebar collapsible="offcanvas">
      <SidebarContent>
        <SidebarGroup class="content-center px-4 w-auto">
          <SidebarGroupLabel class="pt-4 pb-6">Фильтрация по дате</SidebarGroupLabel>
          <DateFilterPicker />
          <SidebarGroupLabel class="pt-8 pb-6">Фитрация по соревнованиям</SidebarGroupLabel>
          <Chooser :options="sports" default-msg="Любой вид спорта" @update="updateSport" />
        </SidebarGroup>
      </SidebarContent>
      <SidebarRail />
    </Sidebar>
    <SidebarInset class="min-h-screen">
      <header
        class="flex w-full h-16 shrink-0 items-center gap-2 transition-[width,height] ease-linear group-has-[[data-collapsible=icon]]/sidebar-wrapper:h-12">
        <div class="flex items-center gap-2 px-4">
          <SidebarTrigger class="-ml-1" />
          <Separator orientation="vertical" class="mr-2 h-4" />
        </div>
      </header>
      <WeeklyView />
    </SidebarInset>
  </SidebarProvider>
</template>
