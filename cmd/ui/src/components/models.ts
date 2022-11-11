export interface Coord {
  x: number;
  y: number;
}

export interface Path {
  coords: Coord[];
}

export interface WordInBoard {
  path: Path;
  bonus?: boolean;
  guessed?: boolean;
}

export interface Board {
  loadedAt?: string;
  width: number;
  height: number;
  cells: string[][];
  words: Record<string, WordInBoard>;
}
