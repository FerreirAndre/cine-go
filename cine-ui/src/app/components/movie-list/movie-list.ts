import { Component, inject, OnInit } from '@angular/core';
import { Movie } from '../../models/movie.model';
import { MovieService } from '../../services/movie.service';
import { Router, RouterLink } from '@angular/router';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-movie-list',
  imports: [RouterLink, CommonModule],
  templateUrl: './movie-list.html',
  styleUrl: './movie-list.css',
})
export class MovieList implements OnInit {
  movies: Movie[] = [];
  loading = true;
  error = false;

  private movieService = inject(MovieService);
  private router = inject(Router);

  ngOnInit(): void {
    this.loadMovies();
  }

  loadMovies() {
    this.movieService.getAll().subscribe({
      next: (data) => {
        this.movies = data;
        this.loading = false;
      },
      error: (err) => {
        this.error = err;
        this.loading = false;
      },
    });
  }

  edit(id: string | undefined) {
    if (id) {
      this.router.navigate(['/movies/edit/', id]);
    }
  }

  delete(id: string | undefined) {
    if (!id) return;

    if (confirm('tem certeza que deseja excluir esse filme?')) {
      this.movieService.delete(id).subscribe({
        next: () => {
          this.movies = this.movies.filter((movie) => movie.ID !== id);
        },
      });
    }
  }

  toggleWatched(id: string | undefined) {
    if (!id) return;

    this.movieService.toggleWatched(id).subscribe({
      next: () => { this.loadMovies(); }
    });
  }
}
