/**
 * Sets a cookie with the given name, value, and optional parameters
 * @param name - The name of the cookie
 * @param value - The value of the cookie
 * @param options - Optional cookie parameters
 */
export function setCookie(
	name: string,
	value: string,
	options: {
		expires?: Date | number;
		path?: string;
		domain?: string;
		secure?: boolean;
		sameSite?: 'strict' | 'lax' | 'none';
	} = {}
): void {
	let cookieString = `${encodeURIComponent(name)}=${encodeURIComponent(value)}`;

	if (options.expires) {
		const expirationDate =
			typeof options.expires === 'number'
				? new Date(Date.now() + options.expires * 24 * 60 * 60 * 1000)
				: options.expires;
		cookieString += `; expires=${expirationDate.toUTCString()}`;
	}

	if (options.path) {
		cookieString += `; path=${options.path}`;
	}

	if (options.domain) {
		cookieString += `; domain=${options.domain}`;
	}

	if (options.secure) {
		cookieString += `; secure`;
	}

	if (options.sameSite) {
		cookieString += `; samesite=${options.sameSite}`;
	}

	document.cookie = cookieString;
}

/**
 * Gets a cookie value by name
 * @param name - The name of the cookie to retrieve
 * @returns The cookie value or null if not found
 */
export function getCookie(name: string, cookie: string): string | null {
	const cookies = cookie.split('; ');
	const encodedName = encodeURIComponent(name);

	for (const cookie of cookies) {
		const [cookieName, cookieValue] = cookie.split('=');

		if (cookieName === encodedName) {
			return decodeURIComponent(cookieValue);
		}
	}

	return null;
}

/**
 * Deletes a cookie by setting its expiration date to the past
 * @param name - The name of the cookie to delete
 * @param options - Optional cookie parameters (path and domain must match the original cookie)
 */
export function deleteCookie(
	name: string,
	options: {
		path?: string;
		domain?: string;
	} = {}
): void {
	setCookie(name, '', {
		...options,
		expires: new Date(0) // Set expiration to epoch time
	});
}
