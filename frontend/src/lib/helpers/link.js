import { PUBLIC_LINKS_URL } from '$env/static/public'

export function linkModel(url) {
  return { 
    "url": url,
    "password": "",
  };
}

export function urlByShortId(shortId) {
  return `${PUBLIC_LINKS_URL}/${shortId}`;
}

export function urlWithoutProtocol(url) {
  return url.replace(/(^\w+:|^)\/\//, '');
}