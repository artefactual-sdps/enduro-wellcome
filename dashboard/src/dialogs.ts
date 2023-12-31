import PackageLocationDialogVue from "./components/PackageLocationDialog.vue";
import { openDialog } from "vue3-promise-dialog";

export async function openPackageLocationDialog(currentLocationId?: string) {
  return await openDialog(PackageLocationDialogVue, {
    currentLocationId: currentLocationId,
  });
}
