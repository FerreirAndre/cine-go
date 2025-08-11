import { Component, inject, Input, OnInit } from '@angular/core';
import {
  FormBuilder,
  FormGroup,
  FormsModule,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { MovieService } from '../../services/movie.service';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-movie-form',
  imports: [ReactiveFormsModule, CommonModule, FormsModule, RouterLink],
  templateUrl: './movie-form.html',
  styleUrl: './movie-form.css',
})
export class MovieForm implements OnInit {
  private fb = inject(FormBuilder);
  private movieService = inject(MovieService);
  private router = inject(Router);
  private route = inject(ActivatedRoute);
  @Input() movieId?: string;

  form!: FormGroup;
  isEditMode = false;

  ngOnInit(): void {
    const idParam = this.route.snapshot.paramMap.get('id');

    if (idParam) {
      this.movieId = idParam;
      this.isEditMode = true;
      this.movieService.getById(idParam).subscribe((movie) => {
        this.form.patchValue(movie);
      });
    }

    this.form = this.fb.group({
      title: ['', Validators.required],
      summary: ['', Validators.required],
      director: ['', Validators.required],
      who_chose: ['', Validators.required],
      release_year: ['', Validators.required],
      cover_link: ['', Validators.required],
      duration: ['', Validators.required],
      rating: ['', Validators.required],
      watched: ['', Validators.required],
    });
  }

  onSubmit() {
    if (this.form.invalid) return;

    if (this.isEditMode && this.movieId) {
      this.movieService.update(this.movieId, this.form.value).subscribe(() => {
        this.router.navigate(['movies']);
      });
    } else {
      this.movieService.create(this.form.value).subscribe(() => {
        this.router.navigate(['movies']);
      });
    }
  }
}
