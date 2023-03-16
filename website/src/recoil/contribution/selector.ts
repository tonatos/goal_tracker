import { selector } from "recoil";
import { contributionsAtom } from "./atom"

export const contributionsListQuery = selector({
    key: 'contributionsListQuery',
    get: ({ get }) => get(contributionsAtom) 
});