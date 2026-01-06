const apiURL = import.meta.env.VITE_API_COSMOS ?? "http://192.168.10.163:1317";
const rpcURL = import.meta.env.VITE_WS_TENDERMINT ?? "http://192.168.10.163:26657";
const prefix = import.meta.env.VITE_ADDRESS_PREFIX ?? "cosmos";

export const env = {
  apiURL,
  rpcURL,
  prefix,
};
