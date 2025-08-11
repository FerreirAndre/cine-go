import { Component, inject, OnInit } from '@angular/core';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';
import { MovieService } from '../../services/movie.service';
import { Movie } from '../../models/movie.model';

@Component({
  selector: 'app-movie-detail',
  imports: [RouterLink],
  templateUrl: './movie-detail.html',
  styleUrl: './movie-detail.css',
})
export class MovieDetail implements OnInit {
  private route = inject(ActivatedRoute);
  private movieService = inject(MovieService);
  private router = inject(Router);

  movie: Movie | null = null;
  loading = true;
  error = false;

  ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.movieService.getById(id).subscribe({
        next: (movie) => {
          this.movie = movie;
          this.loading = false;
        },
        error: () => {
          this.error = true;
          this.loading = false;
        },
      });
    }
  }

  edit(id: string | undefined) {
    if (id) {
      this.router.navigate(['/movies/edit/', id]);
    }
  }

  delete(id: string | undefined) {
    if (!id) return;

    if (confirm('Tem certeza que deseja excluir esse filme?')) {
      this.movieService.delete(id).subscribe({
        next: () => {
          this.router.navigate(['/movies']);
        },
        error: (err) => {
          console.error('erro excluindo filme: ', err);
          alert('erro ao excluir o filme.');
        },
      });
    }
  }
}
