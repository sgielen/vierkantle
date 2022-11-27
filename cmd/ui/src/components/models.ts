export interface Coord {
  x: number;
  y: number;
}

export interface WordInBoard {
  path: Coord[];
  bonus?: boolean;
  guessed?: boolean;
}

export interface Board {
  loadedAt?: string;
  width: number;
  height: number;
  cells: string[][]; // [y][x]
  words: Record<string, WordInBoard>;
}

export interface CellState {
  begins: number;
  used: number;
}
