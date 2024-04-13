import { PUBLIC_API_URL } from '$env/static/public'

export async function fetchLink(link, svelteFetch = undefined) {
  svelteFetch = svelteFetch || fetch;
  return svelteFetch(`${PUBLIC_API_URL}/links/${link}`).then(res => res.json());
}