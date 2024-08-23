export interface CatalogueBook {
  api_id: string;
  title: string;
  subtitle: string;
  authors: string[] | null;
  publishDate: string;
  description: string;
  pageCount: number;
  publisher: string;
  language: string;
  smallThumbnail: string;
  thumbnail: string;
}
