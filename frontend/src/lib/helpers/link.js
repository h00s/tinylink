import { PUBLIC_LINKS_URL } from '$env/static/public'

export function linkModel(url, password = '') {
  return { 
    'url': url,
    'password': password,
  };
}

export function urlByShortId(shortId) {
  return `${PUBLIC_LINKS_URL}/${shortId}`;
}

export function urlWithoutProtocol(url) {
  return url.replace(/(^\w+:|^)\/\//, '');
}