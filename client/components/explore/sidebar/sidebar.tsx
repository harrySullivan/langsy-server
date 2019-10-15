import { h } from 'preact'
import { LensItem } from './lens-item'
import { SidebarItem } from './sidebar-item'
import { SidebarHeading } from './sidebar-heading'
import "./sidebar.css"

export const Sidebar = (
  () => (
    <div class="column pl3">
    	{SidebarHeading("songs")}
    	{LensItem("Songs", "music")}
    	{LensItem("Albums", "compact-disc")}
    	{LensItem("Artists", "microphone")}
    	{LensItem("Liked", "heart")}

        <div class="mt2"></div>
    	{SidebarHeading("your playlists")}
        <div class="mt1"></div>
    	{SidebarItem("ASHWAT")}
    	{SidebarItem("IS")}
    	{SidebarItem("A")}
        {SidebarItem("VERY")}
        {SidebarItem("VERY")}
        {SidebarItem("VERY")}
        {SidebarItem("VERY")}
        {SidebarItem("VERY")}
        {SidebarItem("VERY")}
        {SidebarItem("VERY")}
        {SidebarItem("VERY")}
        {SidebarItem("GAY")}
    	{SidebarItem("PEDOPHILE")}
        <button class="button is-primary is-rounded mt1">+ Create New</button>
    </div>
  )
)