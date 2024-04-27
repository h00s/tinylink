import { fetchLink } from '$lib/repositories/links';
import { redirect } from '@sveltejs/kit';

export async function load({ fetch, params }) {
  const response = await fetchLink(params.link, fetch);
  const data = await response.json();

  return {
    error: data,
  }
}