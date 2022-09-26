import App from "./App.vue";
import auth from "./auth";
import { client, api } from "./client";
import router from "./router";
import { useLayoutStore } from "./stores/layout";
import "./styles/main.scss";
import { PiniaDebounce } from "@pinia/plugin-debounce";
import humanizeDuration from "humanize-duration";
import moment from "moment";
import { createPinia } from "pinia";
import { debounce } from "ts-debounce";
import { createApp } from "vue";
import { PromiseDialog } from "vue3-promise-dialog";

const pinia = createPinia();
pinia.use(PiniaDebounce(debounce));

const app = createApp(App);
app.use(router);
app.use(pinia);
app.use(PromiseDialog);
app.mount("#app");

auth.getUser().then((user) => {
  useLayoutStore().setUser(user);
  if (user) {
    client.package.packageMonitorRequest().then(() => {
      client.connectPackageMonitor();
    });
  }
});

interface Filters {
  [key: string]: (...value: any[]) => string;
}

declare module "@vue/runtime-core" {
  interface ComponentCustomProperties {
    $filters: Filters;
  }
}

app.config.globalProperties.$filters = {
  formatDateTimeString(value: string) {
    const date = new Date(value);
    return date.toLocaleString();
  },
  formatDateTime(value: Date | undefined) {
    if (!value) {
      return "";
    }
    return value.toLocaleString();
  },
  formatDuration(from: Date, to: Date) {
    const diff = moment(to).diff(from);
    return humanizeDuration(moment.duration(diff).asMilliseconds());
  },
  getPreservationActionLabel(
    value: api.EnduroPackagePreservationActionTypeEnum
  ) {
    switch (value) {
      case api.EnduroPackagePreservationActionTypeEnum.CreateAip:
        return "Create AIP";
      case api.EnduroPackagePreservationActionTypeEnum.CreateAndReviewAip:
        return "Create and Review AIP";
      case api.EnduroPackagePreservationActionTypeEnum.MovePackage:
        return "Move package";
      default:
        return value;
    }
  },
  getLocationSourceLabel(value: api.LocationSourceEnum) {
    switch (value) {
      case api.LocationSourceEnum.Minio:
        return "MinIO";
      case api.LocationSourceEnum.Unspecified:
        return "Unspecified";
      default:
        return value;
    }
  },
  getLocationPurposeLabel(value: api.LocationPurposeEnum) {
    switch (value) {
      case api.LocationPurposeEnum.AipStore:
        return "AIP Store";
      case api.LocationPurposeEnum.Unspecified:
        return "Unspecified";
      default:
        return value;
    }
  },
};
