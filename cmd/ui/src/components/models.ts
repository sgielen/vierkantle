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
  name?: string;
  width: number;
  height: number;
  cells: string[][]; // [y][x]
  words: Record<string, WordInBoard>;
  madeBy?: string;
  description?: string;
}

export interface CellState {
  begins: number;
  used: number;
}
