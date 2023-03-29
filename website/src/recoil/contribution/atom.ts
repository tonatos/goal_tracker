import { atom } from "recoil"
import { Contribution } from "../../interfaces/contribution.interfaces";

export const contributionsAtom = atom<Contribution[]>({
    key: "contributionsAtom",
    default: <Contribution[]>[]
});