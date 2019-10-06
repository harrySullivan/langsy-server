import { h } from 'preact'
import { LensItem } from './lens-item.sidebar'
import { SidebarItem } from './sidebar-item.sidebar'
import { SidebarHeading } from './sidebar-heading.sidebar'

export const Sidebar = (
  () => (
    <div class="column pt2 pl3">
    	{SidebarHeading("songs")}
    	{LensItem("Songs", "music")}
    	{LensItem("Albums", "compact-disc")}
    	{LensItem("Artists", "microphone")}
    	{LensItem("Liked", "heart")}

    	{SidebarHeading("your playlists")}
    	{SidebarItem("LDS")}
    	{SidebarItem("WEED")}
    	{SidebarItem("COCAINE")}
    	{SidebarItem("ASHWAT")}
    </div>
  )
)