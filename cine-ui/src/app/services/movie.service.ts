import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Movie } from '../models/movie.model';
import { Observable } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class MovieService {
  private apiUrl = 'http://localhost:8080/movies/';

  constructor(private http: HttpClient) { }

  getAll(): Observable<Movie[]> {
    return this.http.get<Movie[]>(this.apiUrl);
  }

  getById(id: string): Observable<Movie> {
    return this.http.get<Movie>(`${this.apiUrl}${id}`);
  }

  create(movie: Movie): Observable<Movie> {
    return this.http.post<Movie>(this.apiUrl, movie);
  }

  update(id: string, movie: Movie): Observable<void> {
    return this.http.put<void>(`${this.apiUrl}${id}`, movie);
  }

  delete(id: string): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}${id}`);
  }

  toggleWatched(id: string): Observable<void> {
    return this.http.patch<void>(`${this.apiUrl}${id}/watched`, null);
  }
}
