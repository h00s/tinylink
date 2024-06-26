import { PUBLIC_API_URL } from '$env/static/public'
import { linkModel } from '$lib/helpers/link'

export async function fetchLink(link, svelteFetch = undefined) {
  svelteFetch = svelteFetch || fetch;
  return svelteFetch(`${PUBLIC_API_URL}/links/${link}`);
}

export async function createLink(url, password, svelteFetch = undefined) {
  svelteFetch = svelteFetch || fetch;
  return svelteFetch(`${PUBLIC_API_URL}/links`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(linkModel(url, password)),
  });
}

export async function authLink(link, password, svelteFetch = undefined) {
  svelteFetch = svelteFetch || fetch;
  return svelteFetch(`${PUBLIC_API_URL}/links/${link}/authorize`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(linkModel(undefined, password)),
  });
}