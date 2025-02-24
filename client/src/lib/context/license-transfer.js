import { getContext, setContext } from 'svelte';

/**
 * @type {Record<string, string>} map
 */
const store = {};

/**
 * Set the license context
 * @param {string} license
 * @returns {string}
 */
export function setLicenceTransference(license) {
	const randomCode = Math.random().toString(36).slice(2, 8).toString();
	store[randomCode] = license;
	return randomCode;
}

/**
 * Get license
 * @param {string} code
 * @returns {string} license
 */
export function getLicenseTransference(code) {
	const license = store[code];
	delete store[code];
	return license;
}
