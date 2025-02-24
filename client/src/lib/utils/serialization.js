/**
 * Serializes non-POJOs
 * @param {any} obj - The object to serialize
 * @returns {any}
 */
export const serializeNonPOJOs = (obj) => {
	return structuredClone(obj);
};
