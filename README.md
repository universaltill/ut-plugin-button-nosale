# ut-plugin-button-nosale

A `button`-type plugin (ADR-0009 naming: `ut-plugin-{type}-{name}`): adds a
**No Sale** button to the Universal Till sale screen (products panel).

Pressing it publishes `button.nosale.pressed` (`plugin_id`, `entry_key`,
`label`) on the plugin event bus, and the bundled `runtime:"wasm"` handler
(ADR-0001) runs in-process to record the drawer-release request on the audit
trail. Driving a physical cash drawer needs a hardware/device plugin
subscribed to the same event.

This repo is the reference sample for the `button` canonical type.

## Release
Tag `v<version>` → CI validates the manifest, builds the WASI module,
packages a universal `tar.gz`, publishes to the marketplace and (dev)
auto-approves.
Secrets: `MARKETPLACE_BASE_URL`, `MARKETPLACE_UPLOAD_TOKEN`.
Vars: `AUTO_APPROVE`, `MARKETPLACE_LISTING_ID` (set after first publish).
