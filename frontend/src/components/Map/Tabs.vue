<script setup lang="ts">
import { ref, computed } from 'vue';
import { Tab, SubTab } from '../../types/MapEditor/Tab';

const props = withDefaults(defineProps<{
    tabs: Tab[]
}>(), { tabs: (): Tab[] => [] });

const emit = defineEmits<{
    tabChange: [tab: Tab, index: number]
    subTabChange: [subTab: SubTab, index: number]
}>();

const activeTabIndex = ref(0);
const activeSubTabIndex = ref(0);

const activeTab = computed(() => props.tabs[activeTabIndex.value]);
const hasSubTabs = computed(() => !!activeTab.value?.subTabs?.length);

function selectTab(tab: Tab, index: number) {
    activeTabIndex.value = index;
    activeSubTabIndex.value = 0;
    emit('tabChange', tab, index);
}

function selectSubTab(subTab: SubTab, index: number) {
    activeSubTabIndex.value = index;
    emit('subTabChange', subTab, index);
}
</script>

<template>
    <div class="tabs">

        <div class="tab-headers">
            <button
                v-for="(tab, index) in tabs"
                :key="index"
                class="tab-header"
                :class="{ active: activeTabIndex === index }"
                @click="selectTab(tab, index)"
            >
                {{ tab.name }}
            </button>
        </div>

        <div v-if="hasSubTabs" class="sub-tab-headers">
            <button
                v-for="(subTab, index) in activeTab.subTabs"
                :key="index"
                class="sub-tab-header"
                :class="{ active: activeSubTabIndex === index }"
                @click="selectSubTab(subTab, index)"
            >
                {{ subTab.name }}
            </button>
        </div>
    </div>
</template>

<style scoped>
.tabs {
    display: flex;
    flex-direction: column;
    width: 100%;
}

.tab-headers {
    display: flex;
    width: 100%;
    background-color: #1e1e2e;
    border-bottom: 2px solid #313244;
}

.tab-header {
    flex: 1;
    padding: 10px 16px;
    background: transparent;
    border: none;
    border-bottom: 2px solid transparent;
    margin-bottom: -2px;
    color: #6c7086;
    font-size: 13px;
    font-weight: 500;
    cursor: pointer;
    transition: color 0.2s, border-color 0.2s, background-color 0.2s;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    letter-spacing: 0.3px;
}

.tab-header:hover {
    color: #cdd6f4;
    background-color: #2a2a3e;
}

.tab-header.active {
    color: #89b4fa;
    border-bottom-color: #89b4fa;
    background-color: #252538;
}

.sub-tab-headers {
    display: flex;
    width: 100%;
    background-color: #181825;
    border-bottom: 1px solid #313244;
}

.sub-tab-header {
    flex: 1;
    padding: 6px 14px;
    background: transparent;
    border: none;
    border-bottom: 2px solid transparent;
    margin-bottom: -1px;
    color: #585b70;
    font-size: 11px;
    font-weight: 500;
    cursor: pointer;
    transition: color 0.2s, border-color 0.2s, background-color 0.2s;
    white-space: nowrap;
    letter-spacing: 0.3px;
}

.sub-tab-header:hover {
    color: #cdd6f4;
    background-color: #1e1e2e;
}

.sub-tab-header.active {
    color: #a6e3a1;
    border-bottom-color: #a6e3a1;
}

.tab-content {
    flex: 1;
    padding: 16px;
    background-color: #252538;
    color: #cdd6f4;
}
</style>