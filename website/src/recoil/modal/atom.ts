import { atom } from "recoil"
import { ModalState } from "../../interfaces/modal.interfaces";

export const modalStateAtom = atom<ModalState>({
    key: "modalStateAtom",
    default: {} as ModalState,
});
