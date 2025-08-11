import { Routes } from '@angular/router';
import { MovieList } from './components/movie-list/movie-list';
import { MovieForm } from './components/movie-form/movie-form';
import { MovieDetail } from './pages/movie-detail/movie-detail';

export const routes: Routes = [
  { path: '', redirectTo: 'movies', pathMatch: 'full' },
  { path: 'movies', component: MovieList },
  { path: 'movies/new', component: MovieForm },
  { path: 'movies/edit/:id', component: MovieForm },
  { path: 'movies/:id', component: MovieDetail },
];
